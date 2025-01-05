
import FriendItem from '@/components/FriendItem.vue';
import Banner from '@/components/banner/index.vue';
import Msg from '@/components/msg/index.vue';
import axios from 'axios';
import { mapState, mapActions } from 'vuex';
export default {
  components:{
    Msg,
    FriendItem,
    Banner,
  },
  data() {
    return {
      selectedChat: {name:'聊天交友群'},
      msg:'',
      msgs: [],
      socket: null,
      search:'',
      friends:[],
      friend:{},
      remark:'',
      apply_friends:[],
      users:[],
      chats:[],
      activeChat: null, // 当前活动的聊天窗口
      chatWindows: {}, // 用于跟踪每个好友的聊天窗口状态
    }
  },
  methods: {
    ...mapActions(['updatePage', 'updateContent']), // 映射 actions 到组件方法
    formatTime(timestamp){
      return timestamp.substring(11, 19)
    },
    shouldShowTimestamp(index) {
      if (index === 0){
        return true; // 第一条消息总是显示时间
      } 
      const currentTimestamp = this.msgs[index].timestamp.substring(11, 16); // 当前消息的时间
      const previousTimestamp = this.msgs[index - 1].timestamp.substring(11, 16); // 前一条消息的时间

      // 将时间转换为分钟
      const [currentHour, currentMinute] = currentTimestamp.split(':').map(Number);
      const [previousHour, previousMinute] = previousTimestamp.split(':').map(Number);

      // 计算时间差
      const currentTotalMinutes = currentHour * 60 + currentMinute;
      const previousTotalMinutes = previousHour * 60 + previousMinute;

      // 如果时间差大于等于 5 分钟，显示时间
      return (currentTotalMinutes - previousTotalMinutes) >= 5;
    },
    changePage(newPage) {
      this.updatePage(newPage); // 调用 action 更新 page
    },
    changeContent(newContent) {
      this.updateContent(newContent); // 调用 action 更新 content
    },
    connectWebSocket() {
      this.socket = new WebSocket(`ws://192.168.1.5:12345/ws?username=${encodeURIComponent(this.username)}`)

      this.socket.onopen = () => {
        console.log('WebSocket 连接成功')
      }
      
      this.socket.onmessage = (msg) => {
        try {
          const data = JSON.parse(msg.data)
           // 提取用户和好友信息
          const user = data.users;
          const friend = data.friend;
          console.log(data.message,data.msgtype)
          console.log(user,friend)
          if (data.msgtype === 'say') {
            const fuser = {
              friend_id:user.user_id,
              friend_name:user.username,
              avatar_url:user.avatar_url,
            }
            this.friendWindow(fuser)
            this.msgs.push({username:user.username,avatar_url:user.avatar_url,
              message:data.message,timestamp:data.timestamp,user_id:user.user_id,
              friend_id:friend.friend_id})
            this.$nextTick(() => this.scrollToBottom()) // 滚动到底部
          } else if(data.msgtype === 'sayUser') {
            console.log(66666666666)
            this.msgs.push({username:user.username,avatar_url:user.avatar_url,
              message:data.message,timestamp:data.timestamp,user_id:user.user_id,
              friend_id:friend.friend_id})
            this.$nextTick(() => this.scrollToBottom()) // 滚动到底部
          }else if(data.id === 'addFriends') {
            console.log(66666)
            if(data.code === 0){
              this.$message({
                message: data.content,
                type: 'success'
              });
              this.applyFriends() // 当前用户的好友申请列表
            }else if (data.message === '5'){
              this.$message({
                message: "有好友申请了!!!",
                type: 'success'
              });
              this.applyFriends() // 当前用户的好友申请列表
            }else{
              this.$message({
                message: data.content,
                type: 'warning'
              });
            }
            this.search = ''
          } else if (data.msgtype === 'agree') {
            // 处理聊天消息
            if (data.message === '6'){
              this.$message({
                message: "好友同意!!!",
                type: 'success'
              });
              this.fetchFriends();  // 当前用户的好友信息
            }
          }else{
            // 处理系统消息
            const data = JSON.parse(msg.data)
            this.msgs.push({ username: data.username, message: data.message ,
              avatar_url:data.avatar_url,timestamp:data.timestamp})
            this.$nextTick(() => this.scrollToBottom()) // 滚动到底部
          }
        } catch (e) {
          // 异常处理优化：记录详细的异常信息
          console.error('消息解析失败:', e)
        }
      }

      this.socket.onclose = () => {
        console.log('WebSocket 连接已关闭')
      }
      },
      sendMessage() {
        if (!this.msg) {
          this.$message({ message: '消息内容不能为空', type: 'warning' })
          return
        }

        if (this.socket && this.socket.readyState === WebSocket.OPEN) {
          const Data = {
            users: this.users,
            friend: this.friend, // 或者使用 friend.friend_id
            message: this.msg,
            msgtype:"say",
          }
          this.socket.send(JSON.stringify(Data))
          this.msg = '' // 清空输入框
          this.$nextTick(() => this.scrollToBottom())
        }
      },
      scrollToBottom() {
        const chatContainer = this.$refs.messages
        if (chatContainer) {
          chatContainer.scrollTop = chatContainer.scrollHeight // 滚动到底部
        }
      },
      addFriends(){
        this.$prompt('备注', '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
          }).then(({ value }) => {
            this.remark = value
            // 判断不能为空
            if (!this.username || !this.search || !this.remark) {
                console.error("所有字段都必须填写！");
                return; // 退出函数，不发送数据
            }
            const friendsData = {
              username: this.username,
              friendname:this.search,
              remark:this.remark,
              msgtype:"addFriends",
            }
            this.socket.send(JSON.stringify(friendsData));
            this.remark = ''
          }).catch(() => {
            this.$message({
              type: 'info',
              message: '取消输入'
            });       
          });      
      },
      async getUserInfo(){
        try {
          const res = await axios.get(`http://192.168.1.5:12345/user_info?username=${this.username}`);
          this.users = res.data.data
        } catch (error) {
          console.error('请求失败:', error);
          // 处理错误
        }
      },
      async agreeFriend(value){
        try {
          const info = {
            username:this.username,
            friendname:value,
            MsgType:'agree',
          }
          const res = await axios.post('http://192.168.1.5:12345/agree', info);
          this.fetchFriends();  // 当前用户的好友信息
          this.applyFriends() // 当前用户的好友申请列表
          if (res.data.msg === 'ok'){
            this.$message({
              message: "同意好友"+value+"!!!",
              type: 'success'
            });
            this.socket.send(JSON.stringify(info))
          }
        } catch (error) {
          console.error('请求失败:', error);
          // 处理错误
        }
      },
      async fetchFriends() {
        try {
          const response = await axios.get(`http://192.168.1.5:12345/homi?username=${this.username}`);
          this.friends = response.data.data; // 假设返回的数据结构与前述一致
        } catch (error) {
          console.error("获取好友信息失败:", error);
        }
      },
      async applyFriends() {
        try {
          const response = await axios.get(`http://192.168.1.5:12345/apply?username=${this.username}`);
          this.apply_friends = response.data.data; // 假设返回的数据结构与前述一致
        } catch (error) {
          console.error("获取好友信息失败:", error);
        }
      },
      getFriendInfo(friend){
        this.friend = friend
      },
      friendWindow(v){
        const fid = v.friend_id
          if (!this.chatWindows[fid]){
            // 如果没有，创建一个新的聊天窗口
            this.chatWindows[fid] = v; // 将好友信息存储在聊天窗口对象中
            this.chats.push(v); // 将好友添加到聊天列表
          }
          // 切换到该好友的聊天窗口
          this.friend = v || this.friend
          this.activeChat = fid; // 设置当前活动聊天窗口的 ID
          this.changePage('msg')
          this.changeContent('chat')
      },
    },
  computed: {
    ...mapState({
      username: state => state.sharedData.username, // 获取 用户名
      page: state => state.pages.page,       // 获取 page
      content: state => state.pages.content, // 获取 content
    }),
    currentChatWindow(){
      return this.chatWindows[this.activeChat]; // activeChat 是当前选中的朋友 ID
    },
    filteredMessages(){
      return this.msgs.filter(msg => {
        return (msg.user_id === this.users.user_id && msg.friend_id === this.activeChat) ||
               (msg.user_id === this.activeChat && msg.friend_id === this.users.user_id);
      });
    },
    fullAvatarUrl() {
      if (this.users.avatar_url) {
        return require(`@/assets/tx/${this.users.avatar_url}`);
      }
      return null // 或者返回一个默认图片路径
    },
  },
  mounted() {
    this.getUserInfo()  //获得当前用户全部信息
    this.connectWebSocket() // 在组件挂载后连接 WebSocket
    this.fetchFriends();  // 当前用户的好友信息
    this.applyFriends() // 当前用户的好友申请列表
  },
  beforeDestroy() {
    if (this.socket) {
      this.socket.close() // 关闭 WebSocket 连接
      this.socket = null
    }
  },
}