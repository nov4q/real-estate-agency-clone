<template>
  <v-main>
    <v-container>
      <v-row>
        <v-col
          v-for="n in flats"
          :key="n"
          cols="4"
        >
          <v-card
            height="200"
            elevation="9"
            shaped
            outlined
            tile>

            <v-card-title>
              {{ n.title }}
            </v-card-title>
            <v-card-subtitle>
              subtitle
            </v-card-subtitle>
            <v-card-text>
            text
            </v-card-text>
            <v-card-actions>
              <v-btn
                elevation="5"
                small
              >
                button1
              </v-btn>
              <v-btn
                elevation="5"
                small
              >
                button2
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </v-main>

</template>

<script>
export default {
  name: 'DefaultLayout',
  data() {
    return {

      flats: [
        {
          title: 'First flat',
          to:'/',

        },
        {
          title: 'second flat',
          to:'/',

        },
        {
          title: 'second flat',
          to:'/',

        },
        {
          title: 'second flat',
          to:'/',

        },
        {
          title: 'second flat',
          to:'/',

        },
        {
          title: 'second flat',
          to:'/',

        },
        {
          title: 'second flat',
          to:'/',

        },



      ],
      miniVariant: false,
      title: 'Vuetify.js',
    }
  },
}
</script>
