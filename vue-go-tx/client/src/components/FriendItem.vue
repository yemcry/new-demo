<template>
 <div class="friend_item">
    <div class="friend_item_l">
      <img :src="avatar">
      <div class="friend_item_l_r">
        <p>{{ name }}</p>
        <p>{{ remark}}</p>
      </div>
    </div>
    <el-button type="success" size="mini" @click="agreeFriend(name)">接受</el-button>
  </div>
 </template>
 
 <script>
 export default {
   props: {
     avatar: {
       type: String,
       default: require('@/assets/tiktok.png'),
     },
     name: {
       type: String,
       required: true,
     },
     remark: {
       type: String,
       required: true,
     },
   },
   methods:{
    agreeFriend(value){
      this.$emit("agreeFriend",value)
    }
   }
 };
 </script>
 
 <style scoped lang="scss">
 img{
      width: 40px;
      min-width: 40px;
      height: 40px;
      min-height: 40px;
    }
 .friend_item{
    display: flex;
    justify-content: space-between;
    align-items: center;
    .friend_item_l{
      display: flex;
      align-items: center;
      gap: 10px;
      .friend_item_l_r{
        line-height: 5px;
        p:nth-child(2){
          font-size: 12px;
          color: #888;
        }
      }
    }
    
  }
 </style>