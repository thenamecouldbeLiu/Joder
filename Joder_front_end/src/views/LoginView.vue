<template>
  <v-layout>
    <v-main>
      <v-container fluid>
        <v-main>
          <v-row>
            <v-col cols="6"> <left-wrapper></left-wrapper></v-col>
            <v-col cols="6">
              <right-wrapper></right-wrapper>
            </v-col>
          </v-row>
        </v-main>
      </v-container>
    </v-main>
  </v-layout>
  <v-footer class="d-flex flex-column">
    <div class="bg-teal d-flex w-100 align-center px-4">
      <strong>Get connected with us on social networks!</strong>

      <v-spacer></v-spacer>

      <v-btn
        v-for="icon in icons"
        :key="icon"
        class="mx-4"
        :icon="icon"
        variant="plain"
        size="small"
      ></v-btn>
    </div>
  </v-footer>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
// Components
import leftWrapper from '../components/loginComponents/loginLeft.vue'
import rightWrapper from '../components/loginComponents/loginRight.vue'
import type { Ref } from 'vue'

export default defineComponent({
  name: 'LoginView',

  components: { leftWrapper, rightWrapper },
  setup() {
    const icons: string[] = ['mdi-facebook', 'mdi-twitter', 'mdi-linkedin', 'mdi-instagram']
    const isMember: Ref<boolean> = ref(false)
    return { icons, isMember }
  }
})
</script>
