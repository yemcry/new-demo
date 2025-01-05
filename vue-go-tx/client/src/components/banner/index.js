import { mapState, mapActions } from 'vuex';
export default {
  props:{
    fullAvatarUrl:{
      type:String,
      require:true,
      default:'1.jpg'
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
  },
  computed: {
    ...mapState({
      page: state => state.pages.page,       // 获取 page
    }),
  }
}