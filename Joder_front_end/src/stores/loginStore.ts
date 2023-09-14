import { ref, type Ref } from 'vue'
import { defineStore } from 'pinia'
import { type UserInfo } from '@/beans/userInfo';
import { getUserInfo } from '@/API/Common';
import {type AxiosResponse} from 'axios'

export const useLoginStore = defineStore('login', () => {
  const userInfo:Ref<UserInfo|null> = ref(null);
  const isLogin : Ref<boolean> = ref(true)
  async function init(): Promise<AxiosResponse<UserInfo>>{
      const userId: string|null = localStorage.getItem("userId");
      const jwt: string|null = localStorage.getItem("JWT");
      const user:AxiosResponse<UserInfo> = await getUserInfo(userId, jwt);
      return user
  }

  init().then((res: AxiosResponse<UserInfo>)=>{
    userInfo.value = res.data;
  }).catch((error: any)=>{
    console.log('cant get user data' + error);
  })


  return { userInfo, isLogin}
})
