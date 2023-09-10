import { ref, computed, type Ref } from 'vue'
import { defineStore } from 'pinia'

export const useLoginStore = defineStore('login', () => {
  const userId:Ref<string> = ref("");
  const checkLoginStatus = ():boolean=>{
    // 用API 確認session是否還可使用
    return true;
  }

  return { userId}
})
