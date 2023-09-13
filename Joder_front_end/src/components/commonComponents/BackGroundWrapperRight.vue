

<template>
  <right-card></right-card>
</template>
  
  <style scoped>
</style>

<script lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import { useLoginStore } from '@/stores/loginStore'
import rightCard from './rightCard.vue'
import { storeToRefs } from 'pinia'
export default {
  components: { rightCard },

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