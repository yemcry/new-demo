<template>
  <div class="banner_box">
      <div class="banner_t">
        <img :src="fullAvatarUrl" />
        <img @click="changePage('msg'),changeContent('chat')" 
        :src="page === 'msg' ? require('@/assets/svg/消息.svg') : 
        require('@/assets/svg/消息2.svg')" />
        <img @click="changePage('homi')" 
        :src="page === 'homi' ? require('@/assets/svg/好友2.svg') :
         require('@/assets/svg/好友.svg')" />
        <img @click="changePage('pyq')" :src="page === 'pyq' ? require('@/assets/svg/朋友圈2.svg') 
        : require('@/assets/svg/朋友圈1.svg')" />
      </div>
      <div class="banner_b">
        <img :src="require('@/assets/svg/设置.svg')"/>
      </div>
    </div>
</template>

<script src="./index.js"></script>

<style lang="scss" scoped src="./index.scss"></style>