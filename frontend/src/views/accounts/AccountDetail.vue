<template>
  <section class="account-detail">
    <header class="columns page-header">
      <div class="column is-10">
        <h1 class="title is-4 mb-2">{{ title }}</h1>
      </div>
      <div class="column has-text-right">
        <b-button @click="goBack" class="mr-2"> Back </b-button>
        <b-button type="is-primary" @click="goToEdit">Edit</b-button>
      </div>
    </header>

    <b-loading :active="loading.account" />

    <b-tabs type="is-boxed" :animated="false" v-model="activeTab" @input="onTab">
      <b-tab-item
        label="Information"
        label-position="on-border"
        value="info"
        icon="information-outline"
      >
        <section class="wrap">
          <div class="columns">
            <div class="column is-6">
              <strong>Name</strong>:
              <span>{{ account.name || '' }}</span>
            </div>
            <div class="column is-6">
              <strong>Short Name</strong>:
              <span>{{ account.shortName || '' }}</span>
            </div>
          </div>
          <div class="columns">
            <div class="column is-6">
              <strong>Account Type</strong>:
              <span>{{ account.accountType || '' }}</span>
            </div>
            <div class="column is-6">
              <strong>Industry</strong>:
              <span>{{ account.industry || '' }}</span>
            </div>
          </div>
          <div class="columns">
            <div class="column is-6">
              <strong>Sic Code</strong>:
              <span>{{ account.sicCode || '' }}</span>
            </div>
            <div class="column is-6">
              <strong>Website</strong>:
              <span>{{ account.website || '' }}</span>
            </div>
          </div>
          <div class="columns" v-if="account.description">
            <div class="column is-12">
              <strong>Description</strong>:
              <p>{{ account.description }}</p>
            </div>
          </div>
        </section>
      </b-tab-item>
    </b-tabs>
  </section>
</template>

<script>
import Vue from 'vue';
import { mapState } from 'vuex';

export default Vue.extend({
  name: 'AccountDetail',

  data() {
    return {
      title: 'Account Detail',
      account: {},
      activeTab: 'info',
    };
  },

  methods: {
    fetchAccount(id) {
      this.$api.getAccount(id).then((response) => {
        this.account = response || {};
        if (response.name) {
          this.title = response.name;
        } else {
          this.title = 'Account Detail';
        }
      });
    },

    onTab(tab) {
      this.activeTab = tab;
    },

    goToEdit() {
      this.$router.push({ name: 'accountEdit', params: { id: this.$route.params.id } });
    },

    goBack() {
      this.$router.push({ name: 'accounts' });
    },
  },

  computed: {
    ...mapState(['loading', 'settings']),
  },

  mounted() {
    const { id } = this.$route.params;
    if (id) {
      this.fetchAccount(id);
    }
  },
});
</script>

<style lang="scss" scoped>
.account-detail {
  .wrap {
    padding: 1.5rem;
  }

  strong {
    display: inline-block;
    min-width: 120px;
    margin-right: 0.5rem;
  }

  span {
    color: #333;
  }
}
</style>
<template>
  <section class="account-detail">
    <header class="columns page-header">
      <div class="column is-10">
        <h1 class="title is-4 mb-2">{{ title }}</h1>
      </div>
      <div class="column has-text-right">
        <b-button @click="goBack" icon-left="arrow-left" class="mr-2"> Back </b-button>
        <b-button type="is-primary" @click="goToEdit">Edit</b-button>
      </div>
    </header>

    <b-loading :active="loading.account" />

    <b-tabs type="is-boxed" :animated="false" v-model="activeTab" @input="onTab">
      <b-tab-item
        label="Information"
        label-position="on-border"
        value="info"
        icon="information-outline"
      >
        <section class="wrap">
          <div class="columns">
            <div class="column is-6">
              <strong>Name</strong>:
              <span>{{ account.name || '' }}</span>
            </div>
            <div class="column is-6">
              <strong>Short Name</strong>:
              <span>{{ account.shortName || '' }}</span>
            </div>
          </div>
          <div class="columns">
            <div class="column is-6">
              <strong>Account Type</strong>:
              <span>{{ account.accountType || '' }}</span>
            </div>
            <div class="column is-6">
              <strong>Industry</strong>:
              <span>{{ account.industry || '' }}</span>
            </div>
          </div>
          <div class="columns">
            <div class="column is-6">
              <strong>Sic Code</strong>:
              <span>{{ account.sicCode || '' }}</span>
            </div>
            <div class="column is-6">
              <strong>Website</strong>:
              <span>{{ account.website || '' }}</span>
            </div>
          </div>
          <div class="columns" v-if="account.description">
            <div class="column is-12">
              <strong>Description</strong>:
              <p>{{ account.description }}</p>
            </div>
          </div>
        </section>
      </b-tab-item>
    </b-tabs>
  </section>
</template>

<script>
import Vue from 'vue';
import { mapState } from 'vuex';

export default Vue.extend({
  name: 'AccountDetail',

  data() {
    return {
      title: 'Account Detail',
      account: {},
      activeTab: 'info',
    };
  },

  methods: {
    fetchAccount(id) {
      this.$api.getAccount(id).then((response) => {
        this.account = response || {};
        if (response.name) {
          this.title = response.name;
        } else {
          this.title = 'Account Detail';
        }
      });
    },

    onTab(tab) {
      this.activeTab = tab;
    },

    goToEdit() {
      this.$router.push({ name: 'accountEdit', params: { id: this.$route.params.id } });
    },

    goBack() {
      this.$router.push({ name: 'accounts' });
    },
  },

  computed: {
    ...mapState(['loading', 'settings']),
  },

  mounted() {
    const { id } = this.$route.params;
    if (id) {
      this.fetchAccount(id);
    }
  },
});
</script>

<style lang="scss" scoped>
.account-detail {
  .wrap {
    padding: 1.5rem;
  }

  strong {
    display: inline-block;
    min-width: 120px;
    margin-right: 0.5rem;
  }

  span {
    color: #333;
  }
}
</style>
<template>
  <section class="account-detail">
    <header class="columns page-header">
      <div class="column is-10">
        <h1 class="title is-4 mb-2">{{ title }}</h1>
      </div>
      <div class="column has-text-right">
        <b-button @click="goBack" icon-left="arrow-left" class="mr-2"> Back </b-button>
        <b-button type="is-primary" @click="goToEdit">Edit</b-button>
      </div>
    </header>

    <b-loading :active="loading.account" />

    <b-tabs type="is-boxed" :animated="false" v-model="activeTab" @input="onTab">
      <b-tab-item
        label="Information"
        label-position="on-border"
        value="info"
        icon="information-outline"
      >
        <section class="wrap">
          <div class="columns">
            <div class="column is-6">
              <strong>Name</strong>:
              <span>{{ account.name || '' }}</span>
            </div>
            <div class="column is-6">
              <strong>Short Name</strong>:
              <span>{{ account.shortName || '' }}</span>
            </div>
          </div>
          <div class="columns">
            <div class="column is-6">
              <strong>Account Type</strong>:
              <span>{{ account.accountType || '' }}</span>
            </div>
            <div class="column is-6">
              <strong>Industry</strong>:
              <span>{{ account.industry || '' }}</span>
            </div>
          </div>
          <div class="columns">
            <div class="column is-6">
              <strong>Sic Code</strong>:
              <span>{{ account.sicCode || '' }}</span>
            </div>
            <div class="column is-6">
              <strong>Website</strong>:
              <span>{{ account.website || '' }}</span>
            </div>
          </div>
          <div class="columns" v-if="account.description">
            <div class="column is-12">
              <strong>Description</strong>:
              <p>{{ account.description }}</p>
            </div>
          </div>
        </section>
      </b-tab-item>
    </b-tabs>
  </section>
</template>

<script>
import Vue from 'vue';
import { mapState } from 'vuex';

export default Vue.extend({
  name: 'AccountDetail',

  data() {
    return {
      title: 'Account Detail',
      account: {},
      activeTab: 'info',
    };
  },

  methods: {
    fetchAccount(id) {
      this.$api.getAccount(id).then((response) => {
        this.account = response || {};
        if (response.name) {
          this.title = response.name;
        } else {
          this.title = 'Account Detail';
        }
      });
    },

    onTab(tab) {
      this.activeTab = tab;
    },

    goToEdit() {
      this.$router.push({ name: 'accountEdit', params: { id: this.$route.params.id } });
    },

    goBack() {
      this.$router.push({ name: 'accounts' });
    },
  },

  computed: {
    ...mapState(['loading', 'settings']),
  },

  mounted() {
    const { id } = this.$route.params;
    if (id) {
      this.fetchAccount(id);
    }
  },
});
</script>

<style lang="scss" scoped>
.account-detail {
  .wrap {
    padding: 1.5rem;
  }

  strong {
    display: inline-block;
    min-width: 120px;
    margin-right: 0.5rem;
  }

  span {
    color: #333;
  }
}
</style>
