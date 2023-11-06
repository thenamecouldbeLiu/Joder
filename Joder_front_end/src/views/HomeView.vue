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
import { type Ref, defineComponent, ref } from 'vue'
import leftWrapper from '../components/commonComponents/BackGroundWrapperLeft.vue'
import rightWrapper from '../components/commonComponents/BackGroundWrapperRight.vue'
import sideBar from '../components/commonComponents/sideBar.vue'
import { useLoginStore } from '../stores/loginStore'
import { getUnmatched, getUserData } from '@/API/UserApi'
import type { MainUser, User } from '@/beans/userInfo'
import { type AxiosResponse } from 'axios'
import type { ApiModel } from '@/beans/apiResponse'
// Components

export default defineComponent({
  name: 'HomeView',
  setup() {
    const userStore = useLoginStore()
    async function init(): Promise<AxiosResponse<ApiModel<MainUser>>> {
      const user: AxiosResponse<ApiModel<MainUser>> = await getUserData()
      return user
    }
    const unmatched: Ref<User[]> = ref([])
    init()
      .then((res: AxiosResponse<ApiModel<MainUser>>) => {
        userStore.userInfo = res.data.resultBody
        userStore.userInfo.unmatchedPaginateNum = 0
        const unmatchedRes: Promise<AxiosResponse<ApiModel<User[]>>> = getUnmatched(0)
        unmatchedRes.then((res: AxiosResponse<ApiModel<User[]>>) => {
          unmatched.value = res.data.resultBody
        })

        console.log(unmatched)
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
