<template>
  <v-sheet :height="200" border elevation="6" rounded> </v-sheet>
</template>

<script>
import { defineComponent } from 'vue'
import { useSideBarStore } from '@/stores/sideBar.ts'

import { storeToRefs } from 'pinia'
export default defineComponent({
  setup() {
    const sideBar = useSideBarStore()
    const { isExpand, doExpand } = storeToRefs(sideBar)

    return { isExpand, doExpand }
  }
})
</script>