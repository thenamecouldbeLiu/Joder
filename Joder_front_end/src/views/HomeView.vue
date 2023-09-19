<template>

  <v-layout>

    <side-bar></side-bar>
    <v-main>
      <v-container fluid>
        <v-main>
          <v-row>
            <v-col cols="4"> <left-wrapper></left-wrapper></v-col>
            <v-col cols="8">
              <right-wrapper></right-wrapper>
            </v-col>
          </v-row>
        </v-main>
      </v-container>
    </v-main>
  </v-layout>

</template>

<script lang="ts">
import { defineComponent } from 'vue'
import leftWrapper from '../components/commonComponents/BackGroundWrapperLeft.vue'
import rightWrapper from '../components/commonComponents/BackGroundWrapperRight.vue'
import sideBar from '../components/commonComponents/sideBar.vue'
import { useLoginStore } from '../stores/loginStore'
import { getUserInfo } from '@/API/Common'
import type { MainUser } from '@/beans/userInfo'
import { type AxiosResponse } from 'axios'
// Components

export default defineComponent({
  name: 'HomeView',
  setup() {
    const userStore = useLoginStore()
    async function init(): Promise<AxiosResponse<MainUser>> {
      const userId: string | null = localStorage.getItem('userId')
      const jwt: string | null = localStorage.getItem('JWT')
      const user: AxiosResponse<MainUser> = await getUserInfo(userId, jwt)
      return user
    }

    init()
      .then((res: AxiosResponse<MainUser>) => {
        userStore.userInfo = res.data
      })
      .catch((error: any) => {
        console.log('cant get user data' + error)
      })
  },
  components: {
    leftWrapper,
    rightWrapper,
    sideBar
  }
})
</script>
