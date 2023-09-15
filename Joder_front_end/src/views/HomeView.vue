<template>

  <v-layout>

    <side-bar></side-bar>
    <v-main>
      <v-container fluid>
        <v-main>
          <v-row>
            <v-col cols="4"> <left-wrapper></left-wrapper></v-col>
            <v-col cols="8">
              <right-wrapper></right-wrapper>
            </v-col>
          </v-row>
        </v-main>
      </v-container>
    </v-main>
  </v-layout>

</template>

<script lang="ts">
import { defineComponent } from 'vue'
import leftWrapper from '../components/commonComponents/BackGroundWrapperLeft.vue'
import rightWrapper from '../components/commonComponents/BackGroundWrapperRight.vue'
import sideBar from '../components/commonComponents/sideBar.vue'
// Components

export default defineComponent({
  name: 'HomeView',

  components: {
    leftWrapper,
    rightWrapper,
    sideBar,
  }
})
</script>
