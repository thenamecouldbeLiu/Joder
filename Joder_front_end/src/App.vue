<script lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import { useLoginStore } from './stores/loginStore'
import { storeToRefs } from 'pinia'
import BackGroundWrapperLeft from './components/commonComponets/BackGroundWrapperLeft.vue'
import BackGroundWrapperRight from './components/commonComponets/BackGroundWrapperRight.vue'
export default {
  components: { BackGroundWrapperLeft, BackGroundWrapperRight },
  setup() {
    const loginStore = useLoginStore()
    const { userInfo, isLogin } = storeToRefs(loginStore)

    // expose to template and other options API hooks
    return {
      userInfo,
      isLogin
    }
  },

  mounted() {}
}
</script>

<template>
  <div class="col" style="background-color: bisque">
    <back-ground-wrapper-left><slot></slot>></back-ground-wrapper-left>
  </div>

  <div class="col" style="background-color: rgb(250, 247, 247)">
    <back-ground-wrapper-right><slot></slot>></back-ground-wrapper-right>
  </div>
</template>

<style scoped>
.col {
  margin: 1%;
}
</style>
