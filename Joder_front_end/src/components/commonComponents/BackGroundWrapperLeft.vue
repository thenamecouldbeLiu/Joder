

<template><div></div></template>

<script lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import { useLoginStore } from '@/stores/loginStore'
import leftCard from './leftCard.vue'
import { storeToRefs } from 'pinia'
export default {
  components: { leftCard },

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
  
  <style scoped>
</style>