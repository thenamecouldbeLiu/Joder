import { ref, type Ref } from 'vue'
import { defineStore } from 'pinia'
import { type User, type MainUser } from '@/beans/userInfo'

export const useLoginStore = defineStore('login', () => {
  const testOtherUser: User = {
    userId: '00002',
    name: 'otherUser',
    picUrl: 'https://randomuser.me/api/portraits/women/85.jpg',
    email: 'testOther@gmail.com',
    selfIntro: 'testOtherIntro',
    tags: [{ content: 'software' }]
  }

  const testUser: MainUser = {
    userId: '00001',
    name: 'mainUser',
    picUrl: 'https://randomuser.me/api/portraits/women/85.jpg',
    email: 'test@gmail.com',
    selfIntro: 'testIntro',
    tags: [{ content: 'hunter' }],
    pickedUser: [],
    matchedUser: []
  }

  const userInfo: Ref<MainUser | null> = ref(testUser)
  const isLogin: boolean = false

  return { userInfo, isLogin, testUser, testOtherUser }
})
