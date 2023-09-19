<template>
  <v-app>
    <router-view> </router-view>
  </v-app>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

import { useTheme } from 'vuetify'

export default defineComponent({
  name: 'App',
  components: {},

  setup() {}
})
</script>
