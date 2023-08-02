<script>
export default {
    name: "FiatButtons",
    props: ['modelValue'],
    emits: ['update:modelValue'],
}
</script>

<template>
    <input type="radio" class="btn-check" name="options-outlined" @input="$emit('update:modelValue', $event.target.value)"
        value="EUR" id="eur-currency-btn-success-outlined" autocomplete="off" checked>
    <label class="btn btn-outline-primary" for="eur-currency-btn-success-outlined">Euro</label>

    <input type="radio" class="btn-check" name="options-outlined" @input="$emit('update:modelValue', $event.target.value)"
        value="USD" id="usd-currency-check-2-outlined" autocomplete="off">
    <label class="btn btn-outline-warning" for="usd-currency-check-2-outlined">Dollar</label>
</template>