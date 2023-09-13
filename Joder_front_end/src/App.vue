<template>
  <v-app>
    <v-container fluid>
      <left-card v-if="isLogin"></left-card>
      <v-main>
        <v-row>
          <v-col cols="11">
            {{ isLogin }}
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
import leftCard from '@/components/commonComponents/leftCard.vue'
import { storeToRefs } from 'pinia'
export default defineComponent({
  name: 'App',
  components: {
    leftCard
  },

  setup() {
    const userStore = useLoginStore()
    const { userInfo, isLogin } = storeToRefs(userStore)

    return { userInfo, isLogin }
  }
})
</script>
