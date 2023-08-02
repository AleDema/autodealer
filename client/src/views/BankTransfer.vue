<template>
  <form class="deposit-form" @submit.prevent="onSubmit">
    <div class="mb-3">
      <input type="radio" class="btn-check" name="exchange-options-outlined" v-model="exchangeName" value="kraken"
        id="kraken-dark-outlined" autocomplete="off" checked required>
      <label class="btn btn-outline-dark" for="kraken-dark-outlined">Kraken</label>
    </div>

    <div class="mb-3">
      <FiatButtons></FiatButtons>
    </div>
    <div class="mb-3">
      <small id="emailHelp" class="form-text text-muted">This convert all USDT to {{ currency }} and transfers the funds
        to
        an international bank account.</small><br>
      <button class="btn btn-outline-success" type="submit" :disabled="loading">
        <span v-if="loading" class="spinner-border spinner-border-sm" role="status" aria-hidden="true"></span>
        <span v-if="loading">Loading...</span><span v-else>Transfer</span>
      </button>
    </div>
  </form>
  <div v-if="loading" class="spinner-border text-success" role="status">
    <span class="sr-only"></span>
  </div>

  <div v-if="error" class="error">{{ error }}</div>
  <div>
    {{ result }}
  </div>

  <div v-if="address" class="content">
    <p>Balance: {{ balance }} {{ symbol }}</p>
  </div>
</template>

<script>
import FiatButtons from "@/components/form/FiatButtons.vue";
export default {
  name: 'BankTransfer',
  data: function () {
    return {
      loading: false,
      currency: "EUR",
      exchangeName: "kraken",
      error: "",
      balance: "",
      result: "",
    }
  },
  components: {
    FiatButtons
  },
  methods: {
    async onSubmit() {
      // reset previous result
      this.address = ""
      this.balance = ""
      this.loading = true

      try {
        const response = await this.$api.get('bank/transfer/currency')
        console.log(response)
        this.handleData(response)
      } catch (error) {
        console.log(error)
      } finally {
        this.loading = false
      }
    },
    handleData(result) {
      const data = result.data
      this.result = data
      this.address = data.address
      this.symbol = data.code
      this.balance = data.balance
      this.exchangeName = data.exchange
    }
  },
}
</script>
