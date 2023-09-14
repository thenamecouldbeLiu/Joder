<template>
  <v-app>
    <router-view name="homeView"> </router-view>
  </v-app>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { useLoginStore } from '@/stores/loginStore'
import { storeToRefs } from 'pinia'
export default defineComponent({
  name: 'App',
  components: {
    //leftCard
  },

  setup() {
    const userStore = useLoginStore()
    const { userInfo, isLogin } = storeToRefs(userStore)

    return { userInfo, isLogin }
  }
})
</script>
