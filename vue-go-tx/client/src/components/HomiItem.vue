<template>
 <div class="homi_item">
    <img :src="avatar">
    <p>{{ name }}</p>
  </div>
</template>

<script>
export default {
  props: {
    avatar: {
      type: String,
      require:true,
      default: require('@/assets/tiktok.png'),
    },
    name: {
      type: String,
      required: true,
    },
  },
  computed: {
    
  },

};
</script>

<style scoped lang="scss">
img{
  width: 35px;
  min-width: 35px;
  height: 35px;
  min-height: 35px;
}
.homi_item{
  display: flex;
  align-items: center;
  transition: background-color 0.3s, transform 0.2s;
  padding: 5px;
  gap: 10px;
}
.homi_item:hover {
  background-color: #f0f8ff; /* 提示用户可以点击 */
}
.homi_item:active {
  transform: scale(0.95); /* 点击时缩小 */
  background-color: #e0e0e0; /* 点击时的背景颜色（微微灰色） */
}
</style>