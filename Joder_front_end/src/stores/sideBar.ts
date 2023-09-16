import { ref, type Ref } from 'vue'
import { defineStore } from 'pinia'
export const useSideBarStore = defineStore('sideBar', () => {
  const isExpand: Ref<boolean> = ref(false)

  return { isExpand }
})
