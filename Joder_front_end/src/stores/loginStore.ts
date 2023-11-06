import { ref, type Ref } from 'vue'
import { defineStore } from 'pinia'
import { type User, type MainUser } from '@/beans/userInfo'

export const useLoginStore = defineStore(
  'login',
  () => {
    const testOtherUser: User = {
      userId: '00002',
      userType: 'hunter',
      name: 'otherUser',
      picUrl: ['https://randomuser.me/api/portraits/women/85.jpg'],
      email: 'testOther@gmail.com',
      tags: [{ content: 'software' }]
    }

    const userInfo: Ref<MainUser | null> = ref({})

    return { userInfo, testOtherUser }
  },
  {
    persist: { key: 'user', storage: sessionStorage, paths: ['userInfo'] }
  }
)
