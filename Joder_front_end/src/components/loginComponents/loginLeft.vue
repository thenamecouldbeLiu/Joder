<template>
  <v-banner color="error"
            icon="$vuetify"
            lines="one"
            text="Left Wrapper">

  </v-banner>
  <v-container fluid
               class="scrollable backgroundImg"> <v-layout fill-height>

      <v-img class="mx-auto"
             height="450"
             src="">

        <template v-slot:error>
          <!-- show when first image is out-->
          <v-img class="mx-auto"
                 height="300"
                 max-width="500"
                 src="https://picsum.photos/500/300?image=232"></v-img>
        </template>
      </v-img>
    </v-layout></v-container>
</template>

<script lang="ts">
export default {
  setup() {},

  mounted() {}
}
</script>
  
<style >
.backgroundImg {
  background: url('https://cdn.vuetifyjs.com/images/parallax/material.jpg');
  background-size: cover;
  height: 80vh;
}
</style>