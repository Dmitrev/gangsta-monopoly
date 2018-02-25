<template>
    <button v-if="myTurn" type="button" :disabled="thrownDice" @click="throwDice">Throw the Dice</button>
</template>
<script>
    export default {
        computed: {
            myTurn() {
                return this.$store.getters.myTurn;
            },
            thrownDice() {
                return this.$store.getters.thrownDice;
            }
        },
        methods: {
            throwDice: function () {
                if (this.thrownDice) {
                    return false;
                }
                let ws = this.$store.getters.ws;
                if (ws === null) {
                    return;
                }
                ws.send(JSON.stringify({
                    type: "throw_dice"
                }));
            }
        }
    }
</script>