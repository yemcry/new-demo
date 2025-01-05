<template>
  <div class="chat-item" @click="$emit('click')">
    <div class="chat_l">
      <img :src="avatar" />
      <div class="chat-info">
        <span class="chat-name">{{ name }}</span>
        <span class="chat-preview">{{ lastMessage }}</span>
      </div>
    </div>
    <span class="chat-time">{{ lastTime }}</span>
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
    lastMessage: {
      type: (String,Number),
      required: true,
    },
    lastTime: {
      type: String,
      required: true,
    },
  },
  computed: {
    
  },

};
</script>

<style scoped lang="scss">
.chat-item:hover {
  background-color: #f0f8ff; /* 提示用户可以点击 */
}
.chat-item:active {
  transform: scale(0.95); /* 点击时缩小 */
  background-color: #e0e0e0; /* 点击时的背景颜色（微微灰色） */
}
.chat-item {
  display: flex;
  justify-content: space-between;
  transition: background-color 0.3s, transform 0.2s;
  padding: 5px;
  min-width: 100%;
  min-height: 40px;
  background-color: white;
  .chat_l{
    display: flex;
    gap: 20px;
  }
  img {
    min-width: 40px;
    width: 40px;
    min-height: 40px;
    height: 40px;
  }

  .chat-info {
    display: flex;
    flex-direction: column;
    gap: 5px;

    .chat-name {
      font-size: 16px;
    }

    .chat-preview {
      font-size: 12px;
      color: #888;
    }
  }

  .chat-time {
    font-size: 12px;
    color: #888;
  }
}
</style>