<template>
  <v-toolbar density="compact">
    <v-app-bar-nav-icon></v-app-bar-nav-icon>

    <v-toolbar-title>YzLiu</v-toolbar-title>

    <v-spacer></v-spacer>

    <v-btn icon>
      <v-icon>mdi-magnify</v-icon>
    </v-btn>

    <v-btn icon>
      <v-icon>mdi-heart</v-icon>
    </v-btn>

    <v-btn icon>
      <v-icon>mdi-dots-vertical</v-icon>
    </v-btn>
  </v-toolbar>

  <v-container fluid>
    <v-main>
      <v-row>
        <v-col cols="5"> <left-wrapper></left-wrapper></v-col>
        <v-col cols="7">
          <right-wrapper></right-wrapper>
        </v-col>
      </v-row>
    </v-main>
  </v-container>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import leftWrapper from '../components/commonComponents/BackGroundWrapperLeft.vue'
import rightWrapper from '../components/commonComponents/BackGroundWrapperRight.vue'
// Components

export default defineComponent({
  name: 'HomeView',

  components: {
    leftWrapper,
    rightWrapper
  }
})
</script>
