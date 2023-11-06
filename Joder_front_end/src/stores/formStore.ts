import { ref, type Ref } from 'vue'
import { defineStore } from 'pinia'
import type { InsertUser } from '@/beans/userInfo'
export const useFormStore = defineStore(
  'formStore',
  () => {
    const form: Ref<InsertUser> = ref({})

    return { form }
  },
  {
    persist: {
      storage: sessionStorage
    }
  }
)
