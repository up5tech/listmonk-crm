<template>
  <section class="accounts">
    <header class="columns page-header">
      <div class="column is-10">
        <h1 class="title is-4 mb-2">Accounts</h1>
        <span v-if="!isNaN(records.total)">({{ records.total }})</span>
      </div>
      <div class="column has-text-right">
        <b-button type="is-primary" @click="showNewForm">Add Account</b-button>
      </div>
    </header>

    <b-table
      :data="records.results"
      :loading="loading.accounts"
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
          <a class="a" href="#" @click.prevent="deleteAccounts" data-cy="btn-delete-lists">
            <b-icon icon="trash-can-outline" size="is-small" /> {{ $t('globals.buttons.delete') }}
          </a>
          <span class="a">
            {{
              $tc('globals.messages.numSelected', numSelectedAccounts, { num: numSelectedAccounts })
            }}
            <span v-if="!bulk.all && records.total > records.perPage">
              &mdash;
              <a href="#" @click.prevent="onSelectAll" data-cy="select-all-accounts">
                {{ $tc('globals.messages.selectAll', records.total, { num: records.total }) }}
              </a>
            </span>
          </span>
        </div>
      </template>

      <b-table-column
        v-slot="props"
        field="name"
        :label="$t('globals.fields.name')"
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
          <a :href="`/accounts/${props.row.id}`" @click.prevent="showEditForm(props.row)">
            {{ props.row.name }}
          </a>
        </div>
      </b-table-column>

      <b-table-column
        v-slot="props"
        field="shortName"
        label="Short Name"
        header-class="cy-short-name"
        sortable
      >
        {{ props.row.shortName }}
      </b-table-column>

      <b-table-column
        v-slot="props"
        field="accountType"
        label="Account Type"
        header-class="cy-accountType"
        sortable
      >
        {{ props.row.accountType }}
      </b-table-column>

      <b-table-column
        v-slot="props"
        field="industry"
        label="Industry"
        header-class="cy-industry"
        sortable
      >
        {{ props.row.industry }}
      </b-table-column>

      <b-table-column
        v-slot="props"
        field="sicCode"
        label="SIC Code"
        header-class="cy-sic_code"
        sortable
      >
        {{ props.row.sicCode }}
      </b-table-column>

      <b-table-column
        v-slot="props"
        field="website"
        label="Website"
        header-class="cy-website"
        sortable
      >
        {{ props.row.website }}
      </b-table-column>

      <b-table-column
        v-slot="props"
        field="createdAt"
        :label="$t('globals.fields.createdAt')"
        header-class="cy-created_at"
        sortable
      >
        {{ props.row.createdAt }}
      </b-table-column>

      <b-table-column v-slot="props" cell-class="actions" align="right">
        <div>
          <router-link
            v-if="$can('accounts:manage')"
            :to="`/accounts/${props.row.id}`"
            data-cy="btn-detail-account"
          >
            <b-tooltip label="Detail" type="is-dark">
              <b-icon icon="file-find-outline" size="is-small" />
            </b-tooltip>
          </router-link>
        </div>
      </b-table-column>
    </b-table>

    <!-- Add / edit form modal -->
    <b-modal
      scroll="keep"
      :width="600"
      :active.sync="isFormVisible"
      :can-cancel="[{ rafte: onFormClose }]"
    >
      <account-form :data="curItem" :is-editing="isEditing" @finished="onFormClose" />
    </b-modal>

    <p v-if="settings['app.cache_slow_queries']" class="has-text-grey">
      *{{ $t('globals.messages.slowQueriesCached') }}
      <a
        href="https://listmonk.app/docs/maintenance/performance/"
        target="_blank"
        rel="noopener noreferer"
        class="has-text-grey"
      >
        <b-icon icon="link-variant" /> {{ $t('globals.buttons.learnMore') }}
      </a>
    </p>
  </section>
</template>

<script>
import Vue from 'vue';
import { mapState } from 'vuex';
import AccountForm from './AccountForm.vue';
// import EmptyPlaceholder from '../../components/EmptyPlaceholder.vue';

export default Vue.extend({
  components: {
    // EmptyPlaceholder,
    AccountForm,
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

      // Table bulk row selection states.
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

    onPageChange(page) {
      this.queryParams.page = page;
      this.fetchRecords();
    },

    onSort(field, direction) {
      this.queryParams.orderBy = field;
      this.queryParams.order = direction;
      this.fetchRecords();
    },

    fetchRecords() {
      this.$api
        .getAccounts({
          page: this.queryParams.page,
          query: this.queryParams.query,
          orderBy: this.queryParams.orderBy,
          order: this.queryParams.order,
        })
        .then((data) => {
          this.records = data;
        });
    },

    // Mark all lists in the query as selected.
    onSelectAll() {
      this.bulk.all = true;
    },

    onTableCheck() {
      // Disable bulk.all selection if there are no rows checked in the table.
      if (this.bulk.checked.length !== this.records.total) {
        this.bulk.all = false;
      }
    },
  },

  computed: {
    ...mapState(['loading', 'settings']),

    numSelectedAccounts() {
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
