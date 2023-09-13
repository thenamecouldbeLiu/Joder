import { ref, type Ref } from 'vue'
import { defineStore } from 'pinia'
import { type UserInfo } from '@/beans/userInfo';
export const useLoginStore = defineStore('login', () => {
  const userInfo:Ref<UserInfo|null> = ref(null);
  const isLogin : Ref<boolean> = ref(true)

  return { userInfo, isLogin}
})
