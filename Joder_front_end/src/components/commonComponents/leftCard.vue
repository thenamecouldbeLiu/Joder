<template>
  <v-card>
    <v-navigation-drawer theme="dark" rail permanent>
      <v-list-item
        nav
        prepend-avatar="https://randomuser.me/api/portraits/women/75.jpg"
      ></v-list-item>

      <v-divider></v-divider>

      <v-list density="compact" nav>
        <v-list-item prepend-icon="mdi-view-dashboard" value="dashboard"></v-list-item>

        <v-list-item prepend-icon="mdi-forum" value="messages"></v-list-item>
      </v-list>
    </v-navigation-drawer>

    <v-navigation-drawer permanent>
      <v-list>
        <v-list-item title="Home" value="home"></v-list-item>

        <v-list-item title="Contacts" value="contacts"></v-list-item>

        <v-list-item title="Settings" value="settings"></v-list-item>
      </v-list>
    </v-navigation-drawer>

    <!--<v-main style="height: 300px"></v-main>-->
  </v-card>
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