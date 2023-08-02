<template>
  <form class="deposit-form" @submit.prevent="onSubmit">

    <div class="input-group mb-3">
      <ExchangeInput v-model="exchangeName" label="Exchange" id="exchange" />
    </div>

    <div class="row mb-2">
      <div class="col">
        <v-select v-model="pair" :options="pairs" label="name" placeholder="BTC-USD"></v-select>
      </div>
      <div class="col">
        <label class="form-check-label">{{ pair.assetType }}</label>
      </div>
    </div>
    <div class="mb-3">
      <div class="row">
        <div class="col">
          <input type="number" class="form-control" v-model="qtyUSD" placeholder="USD QTY" step="any" aria-label="USD">
        </div>
        <!--              <div class="col">-->
        <!--                <input type="number" class="form-control" v-model="sizeQTY" placeholder="qty" aria-label="QTY">-->
        <!--              </div>-->
      </div>
    </div>

    <div class="mb-3">
      <div class="row">
        <div class="col">
          <!-- //order type market -->
          <div class="row">
            <div class="col">
              <input type="radio" class="btn-check" name="orderTypeMarketBuy-options-outlined" v-model="orderType"
                value="marketBuy" id="orderTypeMarketBuy-info-outlined" autocomplete="off">
              <label class="btn btn-outline-info" for="orderTypeMarketBuy-info-outlined">Market buy</label>
            </div>
            <div class="col">
              <input type="radio" class="btn-check" name="orderTypeMarketBuy-options-outlined" v-model="orderType"
                value="marketSell" id="orderTypeMarketSell-danger-outlined" autocomplete="off">
              <label class="btn btn-outline-danger" for="orderTypeMarketSell-danger-outlined">Market sell</label>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="mb-3">
      <div class="row">
        <div class="col">
          <!-- //"orderTypeLimit" -->
          <div class="row mb-3">
            <div class="col">
              <input type="radio" class="btn-check" name="orderTypeLimitBuy-options-outlined" v-model="orderType"
                value="limitBuy" id="orderTypeLimitBuy-info-outlined" autocomplete="off">
              <label class="btn btn-outline-info" for="orderTypeLimitBuy-info-outlined">Limit buy</label>
            </div>
            <div class="col">
              <input type="radio" class="btn-check" name="orderTypeLimitSell-options-outlined" v-model="orderType"
                value="limitSell" id="orderTypeLimitSell-danger-outlined" autocomplete="off">
              <label class="btn btn-outline-danger" for="orderTypeLimitSell-danger-outlined">Limit sell</label>
            </div>
          </div>
        </div>
      </div>

      <div class="mb-3">
        <div class="row">
          <div class="col">
            <div class="form-check form-switch">
              <input class="form-check-input" type="checkbox" role="switch" id="reduceFlexSwitchCheckDefault">
              <label class="form-check-label" for="reduceFlexSwitchCheckDefault">Reduce</label>
            </div>
          </div>

          <div class="col">
            <div class="form-check form-switch">
              <input class="form-check-input" type="checkbox" role="switch" v-model="twap" id="flexSwitchCheckDefault">
              <label class="form-check-label" for="reduceFlexSwitchCheckDefault">TWAP</label>
            </div>
          </div>
        </div>
      </div>

      <div class="mb-3" v-if="twap">
        <div class="row">
          <div class="col">
            <input type="number" class="form-control" placeholder="hours" v-model="hours" aria-label="hours"
              aria-describedby="basic-addon1">
          </div>
          <div class="col">
            <input type="number" class="form-control" placeholder="minutes" v-model="minutes" aria-label="minutes"
              aria-describedby="basic-addon2">
          </div>
        </div>
      </div>
    </div>

    <button class="btn btn-primary" type="submit" :disabled="loading">
      <span v-if="loading" class="spinner-border spinner-border-sm" role="status" aria-hidden="true"></span>
      <span v-if="loading">Loading...</span><span v-else>Submit</span>
    </button>
  </form>

  <div class="mt-2">{{ result }}</div>
  <Chart :options="{
    width: 1200,
    height: 675,
    symbol: 'BINANCE:BTCUSD',
    interval: 'D',
    timezone: 'Europe/Amsterdam',
    theme: 'dark',
    style: '1',
    locale: 'en',
    toolbar_bg: '#f1f3f6',
    enable_publishing: false,
    withdateranges: true,
    hide_side_toolbar: false,
    allow_symbol_change: true,
    details: true,
    studies: [
      'VSTOP@tv-basicstudies',
      'Volume@tv-basicstudies',
      'VWAP@tv-basicstudies'
    ],
    show_popup_button: true,
    popup_width: '800',
    popup_height: '450',
  }" />
</template>

<script>
import { Chart } from 'vue-tradingview-widgets';
import vSelect from 'vue-select'
import ExchangeInput from "@/components/form/ExchangeInput.vue";
export default {
  name: 'Trade',
  data: function () {
    return {
      pair: { assetType: 'spot', name: 'BTC' },
      pairs: [],
      loading: false,
      error: "",
      exchangeName: "ftx",
      qtyUSD: "",
      orderType: "",
      side: "",
      type: "",
      balance: "",
      assetType: "spot",
      twap: false,
      hours: 0,
      minutes: 0,
      notification: 0,
      result: "",
    }
  },
  components: {
    Chart,
    vSelect,
    ExchangeInput
  },
  mounted() {
    this.onPairsLoaded();
  },
  methods: {
    onPairsLoaded: function () {
      this.pairs = []; // clear the array
      this.$api.get('pairs/' + this.exchangeName).then(response => {
        this.pairs = response.data["pair"];
      }).catch(error => {
        console.log(error);
      });

      try {
        const response = this.$api.get('pairs/' + this.exchangeName)
        this.pairs = response.data["pair"];
      } catch (error) {
        console.log(error)
      }
    },
    // onOrderType switch case for order type
    onOrderType: function () {
      switch (this.orderType) {
        case "limitBuy":
          this.side = "buy"
          this.orderType = "limit"
          break;
        case "marketBuy":
          this.side = "buy"
          this.orderType = "market"
          break;
        case "limitSell":
          this.side = "sell"
          this.orderType = "limit"
          break;
        case "marketSell":
          this.side = "sell"
          this.orderType = "market"
          break;
        default:
          this.side = "buy"
          this.orderType = "market"
          break;
      }
    },
    // axios call to place a TWAP order on the exchange
    onTWAPOrder: async function () {
      try {
        const response = this.$api.get('twap/' + this.exchangeName + '/' + this.pair.name + '/' + this.qtyUSD + '/' + this.pair.assetType + '/' + this.orderType + '/' + this.side + '/' + this.hours + '/' + this.minutes)

        this.processData(response)
      } catch (error) {
        console.log(error)
        this.error = error
      } finally {
        this.loading = false
      }
    },
    onTradeOrder: async function () {
      try {
        const response = this.$api.get('trade/' + this.exchangeName + '/' + this.pair.name + '/' + this.qtyUSD + '/' + this.pair.assetType + '/' + this.orderType + '/' + this.side)

        this.processData(response)
      } catch (error) {
        console.log(error)
        this.error = error
      } finally {
        this.loading = false
      }
    },
    onSubmit() {
      // reset previous result
      this.result = ""
      this.error = ""

      // reset address and balance for new result
      this.balance = ""

      // disable the form
      this.loading = true

      this.onOrderType()

      if (this.twap) {
        this.onTWAPOrder()
      } else {
        this.onTradeOrder()
      }
    },

    // axios call to create a new TWAP order
    processData(result) {
      const data = result.data
      this.result = data
    }
  },
  watch: {
    exchangeName: {
      handler: function (value) {
        this.onPairsLoaded();
      }
    }
  }
}
</script>
