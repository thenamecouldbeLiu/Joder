<template>
  <v-app>
    <v-toolbar density="compact">
      <v-app-bar-nav-icon></v-app-bar-nav-icon>

      <v-toolbar-title>YzLiu</v-toolbar-title>

      <v-spacer></v-spacer>

      <v-btn icon>
        <v-icon>mdi-magnify</v-icon>
      </v-btn>

      <v-btn icon>
        <v-icon>mdi-heart</v-icon>
      </v-btn>

      <v-btn icon>
        <v-icon>mdi-dots-vertical</v-icon>
      </v-btn>
    </v-toolbar>

    <v-container fluid>
      <v-main>
        <v-row>
          <v-col cols="5"> <router-view name="leftWrapper"> </router-view></v-col>
          <v-col cols="7">
            <router-view name="rightWrapper"> </router-view>
          </v-col>
        </v-row>
      </v-main>
    </v-container>
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
