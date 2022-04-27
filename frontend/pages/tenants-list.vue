<template>
 <v-card class="mx-auto" tile>
  <v-row>
    <v-col >

      <v-list flat >
        <v-subheader>Tenants</v-subheader>
        <v-list-item-group
          v-model="selectedItem"
          color="primary"
        >

        <v-list-item
          v-for="(item, i) in tenants"
          :key="i"
          :to="item.to"
          router
          exact
        >

            <v-list-item-content>
              <v-list-item-title>
                {{ item.title }}
              </v-list-item-title>

              <v-list-item-subtitle>
                {{ item.secondLine }}
              </v-list-item-subtitle>

              <v-list-item-subtitle>
                {{ item.thirdLine }}
              </v-list-item-subtitle>

            </v-list-item-content>
          </v-list-item>
        </v-list-item-group>
      </v-list>

    </v-col>
  </v-row>
 </v-card>
</template>

<script>
export default {
  name: 'DefaultLayout',
  data() {
    return {

      tenants: [
        {
          title: 'Tenant',
          secondLine: 'first line of informations',
          thirdLine:'second line of informations',
        },
        {
          title: 'Tenant',
          secondLine: 'Secondary line text Lorem ipsum dolor sit amet,',
          thirdLine:'consectetur adipiscing elit.',
        },
        {
          title: 'Tenant',
          secondLine: 'Secondary line text Lorem ipsum dolor sit amet,',
          thirdLine:'consectetur adipiscing elit.',
        },


      ],
      miniVariant: false,
      title: 'Vuetify.js',
    }
  },
}
</script>

