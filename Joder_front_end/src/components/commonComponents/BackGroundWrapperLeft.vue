

<template>
  <v-banner color="error"
            icon="$vuetify"
            lines="one"
            text="Potential Candidates">

  </v-banner>
  <v-container fluid
               class=""> <v-layout fill-height><left-card /></v-layout></v-container>
</template>

<script lang="ts">
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
.scrollable {
  height: 80vh;
  overflow-x: hidden;
  overflow-y: auto;
  text-align: center;
}
</style>