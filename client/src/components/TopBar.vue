<template>
  <v-toolbar
    absolute
    app
    clipped-left
    dense
    flat
    dark
    class="indigo darken-1"
  >
    <v-toolbar-side-icon
      @click.stop="toggleDrawer"
    ></v-toolbar-side-icon>
    <v-toolbar-title><router-link to="/dashboard" style="color: #fff; text-decoration: none;">Portals@me</router-link></v-toolbar-title>

    <v-spacer></v-spacer>

    <v-chip color="purple lighten-2" small v-if="$config.isDev">
      Development
    </v-chip>

    <v-toolbar-items v-if="user != null">
      <v-menu offset-y>
        <v-btn
          slot="activator"
          flat
        >
          <v-avatar color="orange" size="32px">
            <v-img :src="user.picture" />
          </v-avatar>
          &nbsp;&nbsp;{{ user.display_name }}
        </v-btn>
        <v-list>
          <v-list-tile
            @click="$router.push(`/users/${user.name}`)"
          >
            <v-list-tile-title>プロフィール</v-list-tile-title>
          </v-list-tile>
          <v-list-tile
            @click="signOut"
          >
            <v-list-tile-title>サインアウト</v-list-tile-title>
          </v-list-tile>
        </v-list>
      </v-menu>
    </v-toolbar-items>
  </v-toolbar>
</template>

<script>
export default {
  data () {
    return {
      user: null,
    };
  },
  methods: {
    async signOut () {
      localStorage.setItem('id_token', '');
      localStorage.setItem('user', '{}');
      this.user = null;
      this.$router.push('/');
    },
    toggleDrawer () {
      this.$store.commit('toggleDrawer')
    }
  },
  async mounted () {
    try {
      this.user = JSON.parse(localStorage.getItem('user'));
    } catch (e) {
      console.error(e);
      this.user = {};
    }

    const needAuth = (str) => str.match(
      /^\/collections\/(.*)/
    );

    if (!this.user && !needAuth(this.$route.path)) {
      this.$router.push('/');
    }
  },
}
</script>
