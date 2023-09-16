import { ref, type Ref } from 'vue'
import { defineStore } from 'pinia'
import { type User, type MainUser } from '@/beans/userInfo'
import { getUserInfo } from '@/API/Common'
import { type AxiosResponse } from 'axios'

export const useLoginStore = defineStore('login', () => {
  const testOtherUser: User = {
    userId: '00002',
    name: 'otherUser',
    picUrl: 'https://randomuser.me/api/portraits/women/85.jpg',
    email: 'testOther@gmail.com',
    selfIntro: 'testOtherIntro',
    tags: [{content:'software'}]
  }

  const testUser: MainUser = {
    userId: '00001',
    name: 'mainUser',
    picUrl: 'https://randomuser.me/api/portraits/women/85.jpg',
    email: 'test@gmail.com',
    selfIntro: 'testIntro',
    tags: [{content:'hunter'}],
    pickedUser: [],
    matchedUser: []
  }

  const userInfo: Ref<MainUser | null> = ref(testUser)
  const isLogin: Ref<boolean> = ref(true)
  async function init(): Promise<AxiosResponse<MainUser>> {
    const userId: string | null = localStorage.getItem('userId')
    const jwt: string | null = localStorage.getItem('JWT')
    const user: AxiosResponse<MainUser> = await getUserInfo(userId, jwt)
    return user
  }

  init()
    .then((res: AxiosResponse<MainUser>) => {
      userInfo.value = res.data
    })
    .catch((error: any) => {
      console.log('cant get user data' + error)
    })

  return { userInfo, isLogin, testUser, testOtherUser }
})
