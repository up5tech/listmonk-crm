<template>
  <form @submit.prevent="onSubmit">
    <div class="modal-card content" style="width: auto">
      <header class="modal-card-head">
        <h4 v-if="isEditing">
          {{ data.name }}
        </h4>
        <h4 v-else>New Account</h4>
      </header>
      <section expanded class="modal-card-body">
        <b-field :label="$t('globals.fields.name')" label-position="on-border">
          <b-input
            :maxlength="200"
            :ref="'focus'"
            v-model="form.name"
            name="name"
            :placeholder="$t('globals.fields.name')"
            required
          />
        </b-field>
        <b-field label="Short Name" label-position="on-border">
          <b-input
            :maxlength="100"
            v-model="form.shortName"
            name="shortName"
            placeholder="Short Name"
          />
        </b-field>
        <b-field label="Account Type" label-position="on-border">
          <b-select v-model="form.accountType" name="accountType">
            <option value="direct">Direct</option>
            <option value="agency">Agency</option>
          </b-select>
        </b-field>
        <b-field label="Industry" label-position="on-border">
          <b-select v-model="form.industry" name="industry">
            <option value="industry">Industry</option>
          </b-select>
        </b-field>
        <b-field label="Sic Code" label-position="on-border">
          <b-input :maxlength="10" v-model="form.sicCode" name="sicCode" placeholder="Sic Code" />
        </b-field>
        <b-field label="Website" label-position="on-border">
          <b-input :maxlength="100" v-model="form.website" name="website" placeholder="Website" />
        </b-field>
        <b-field :label="$t('globals.fields.description')" label-position="on-border">
          <b-input
            :maxlength="2000"
            v-model="form.description"
            name="description"
            :placeholder="$t('globals.fields.description')"
            type="textarea"
          />
        </b-field>
      </section>
      <footer class="modal-card-foot">
        <b-button @click="$parent.close()">
          {{ $t('globals.buttons.close') }}
        </b-button>
        <b-button
          v-if="$can('accounts:manager') || $canList(data.id, 'accounts:manage')"
          native-type="submit"
          type="is-primary"
          :loading="loading.accounts"
          data-cy="btn-save"
        >
          {{ $t('globals.buttons.save') }}
        </b-button>
      </footer>
    </div>
  </form>
</template>

<script>
import Vue from 'vue';
import { mapState } from 'vuex';

export default Vue.extend({
  props: {
    data: { type: Object, default: () => ({}) },
    isEditing: { type: Boolean, default: false },
  },

  data() {
    return {
      form: {
        name: '',
        shortName: '',
        accountType: '',
        industry: '',
        sicCode: '',
        website: '',
        description: '',
      },
    };
  },

  computed: {
    ...mapState(['loading', 'profile']),
  },

  mounted() {
    this.form = { ...this.form, ...this.$props.data };
  },

  methods: {
    onSubmit() {
      if (this.isEditing) {
        this.updateAccount();
      } else {
        this.createAccount();
      }
    },

    createAccount() {
      this.$api.createAccount(this.form).then((data) => {
        this.$emit('finished');
        this.$parent.close();
        this.$utils.toast(this.$t('globals.messages.created', { name: data.name }));
      });
    },

    updateAccount() {
      this.$api.updateAccount(this.form).then((data) => {
        this.$emit('finished');
        this.$parent.close();
        this.$utils.toast(this.$t('globals.messages.updated', { name: data.name }));
      });
    },
  },
});
</script>
