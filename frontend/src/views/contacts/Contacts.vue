<template>
  <section class="contacts">
    <header class="columns page-header">
      <div class="column is-10">
        <h1 class="title is-4 mb-2">Contacts</h1>
        <span v-if="!isNaN(records.total)">({{ records.total }})</span>
      </div>
      <div class="column has-text-right">
        <b-button type="is-primary" @click="showNewForm">Add Contact</b-button>
      </div>
    </header>

    <b-table
      :data="records.results"
      :loading="loading.contacts"
      @check-all="onTableCheck"
      @check="onTableCheck"
      :checked-rows.sync="bulk.checked"
      hoverable
      default-sort="createdAt"
      paginated
      backend-pagination
      pagination-position="both"
      @page-change="onPageChange"
      :current-page="queryParams.page"
      :per-page="records.perPage"
      :total="records.total"
      checkable
      backend-sorting
      @sort="onSort"
    >
      <template #top-left>
        <div class="columns">
          <div class="column is-6">
            <form @submit.prevent="fetchRecords">
              <b-field>
                <b-input
                  v-model="queryParams.query"
                  name="query"
                  expanded
                  icon="magnify"
                  ref="query"
                  data-cy="query"
                />
                <p class="controls">
                  <b-button
                    native-type="submit"
                    type="is-primary"
                    icon-left="magnify"
                    data-cy="btn-query"
                  />
                </p>
              </b-field>
            </form>
          </div>
        </div>
        <div class="actions" v-if="bulk.checked.length > 0">
          <a class="a" href="#" @click.prevent="deleteContacts" data-cy="btn-delete-lists">
            <b-icon icon="trash-can-outline" size="is-small" /> {{ $t('globals.buttons.delete') }}
          </a>
          <span class="a">
            {{
              $tc('globals.messages.numSelected', numSelectedContacts, { num: numSelectedContacts })
            }}
            <span v-if="!bulk.all && records.total > records.perPage">
              &mdash;
              <a href="#" @click.prevent="onSelectAll" data-cy="select-all-contacts">
                {{ $tc('globals.messages.selectAll', records.total, { num: records.total }) }}
              </a>
            </span>
          </span>
        </div>
      </template>

      <b-table-column
        v-slot="props"
        field="name"
        label="First Name"
        header-class="cy-name"
        sortable
        width="25%"
        paginated
        backend-pagination
        pagination-position="both"
        :td-attrs="$utils.tdID"
        @page-change="onPageChange"
      >
        <div>
          <a :href="`/contacts/${props.row.id}`" @click.prevent="showEditForm(props.row)">
            {{ props.row.firstName }}
          </a>
        </div>
      </b-table-column>

      <b-table-column
        v-slot="props"
        field="last_name"
        label="Last Name"
        header-class="cy-last-name"
        sortable
        width="25%"
      >
        {{ props.row.lastName }}
      </b-table-column>

      <b-table-column
        v-slot="props"
        field="email"
        label="Email"
        header-class="cy-email"
        sortable
        width="30%"
      >
        {{ props.row.email }}
      </b-table-column>

      <b-table-column
        v-slot="props"
        field="phone"
        label="Phone"
        header-class="cy-phone"
        sortable
        width="20%"
      >
        {{ props.row.phone }}
      </b-table-column>
    </b-table>

    <b-modal
      scroll="keep"
      :width="600"
      :active.sync="isFormVisible"
      :can-cancel="[{ rafte: onFormClose }]"
    >
      <contact-form :data="curItem" :is-editing="isEditing" @finished="onFormClose" />
    </b-modal>
  </section>
</template>
<script>
import Vue from 'vue';
import { mapState } from 'vuex';
import ContactForm from './ContactForm.vue';

export default Vue.extend({
  name: 'Contacts',

  components: {
    ContactForm,
  },

  data() {
    return {
      curItem: null,
      isEditing: false,
      isFormVisible: false,
      records: {},
      queryParams: {
        page: 1,
        perPage: 20,
        query: '',
        orderBy: 'id',
        order: 'asc',
      },
      bulk: {
        checked: [],
        all: false,
      },
    };
  },

  methods: {
    showNewForm() {
      this.curItem = {};
      this.isFormVisible = true;
      this.isEditing = false;
    },

    showEditForm(item) {
      this.curItem = item;
      this.isFormVisible = true;
      this.isEditing = true;
    },

    onFormClose() {
      this.curItem = null;
      this.isFormVisible = false;
      this.isEditing = false;
      this.fetchRecords();
    },

    onTableCheck(selection) {
      this.bulk.checked = selection;
    },

    onSelectAll() {
      this.bulk.all = true;
      this.bulk.checked = [];
    },

    fetchRecords() {
      this.bulk.checked = [];
      this.bulk.all = false;
      this.$api.getContacts(this.queryParams).then((response) => {
        this.records = response || {};
      });
    },

    onPageChange(page) {
      this.queryParams.page = page;
      this.fetchRecords();
    },

    onSort(field, order) {
      this.queryParams.orderBy = field;
      this.queryParams.order = order;
      this.fetchRecords();
    },

    async deleteContacts() {
      const confirmText = this.$t('globals.messages.confirmDelete', {
        num: this.numSelectedContacts,
      });
      const ok = await this.$buefy.dialog.confirm({
        message: confirmText,
        confirmText: this.$t('globals.buttons.delete'),
        cancelText: this.$t('globals.buttons.cancel'),
        type: 'is-danger',
      });
      if (!ok) {
        return;
      }

      this.$store.dispatch('setLoading', { key: 'contacts', value: true });
      try {
        const ids = this.bulk.all
          ? this.records.results.map((r) => r.id)
          : this.bulk.checked.map((r) => r.id);
        await this.$api.deleteContacts(ids);
        this.$buefy.toast.success(this.$t('globals.messages.deleted'));
        this.fetchRecords();
      } finally {
        this.$store.dispatch('setLoading', { key: 'contacts', value: false });
      }
    },
  },

  computed: {
    ...mapState(['loading', 'settings']),
    numSelectedContacts() {
      return this.bulk.all ? this.records.total : this.bulk.checked.length;
    },
  },

  created() {
    this.$root.$on('page.refresh', this.fetchRecords);
  },

  destroyed() {
    this.$root.$off('page.refresh', this.fetchRecords);
  },

  mounted() {
    this.fetchRecords();
  },
});
</script>
