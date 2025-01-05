import ChatItem from '@/components/ChatItem.vue';
import HomiItem from '@/components/HomiItem.vue';
import { mapState, mapActions } from 'vuex';
export default{
  props:{
    friends:{
      type:Array,
      require:true
    },
    chats:{
      type:Array,
      require:true
    },
  },
  components:{
    ChatItem,
    HomiItem,
  },
  data() {
    return {
      lastTime:'现在'
    }
  },
  methods:{
    ...mapActions(['updatePage', 'updateContent']), // 映射 actions 到组件方法
    changePage(newPage) {
      this.updatePage(newPage); // 调用 action 更新 page
    },
    changeContent(newContent) {
      this.updateContent(newContent); // 调用 action 更新 content
    },
    handleChatSelected(chat) {
      this.selectedChat = chat
    },
    handleClick(friend) {
      this.changeContent('friend_info');
      this.getFriendInfo(friend);
    },
    getFriendInfo(friend){
      this.$emit("getFriendInfo",friend)
    },
    friendWindow(value){
      this.$emit("friendWindow",value)
    },
  },
  computed: {
    ...mapState({
      page: state => state.pages.page,       // 获取 page
      content: state => state.pages.content, // 获取 content
    }),
  },
}