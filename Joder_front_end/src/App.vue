<template>
  <v-app>
    <router-view> </router-view>
  </v-app>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { useLoginStore } from './stores/loginStore'
import { storeToRefs } from 'pinia'
import { useTheme } from 'vuetify'

export default defineComponent({
  name: 'App',
  components: {},

  setup() {
    const userStore = useLoginStore()
    const { userInfo, isLogin } = storeToRefs(userStore)

    return { userInfo, isLogin }
  }
})
</script>
