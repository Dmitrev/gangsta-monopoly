<template>
    <button v-if="myTurn" type="button" :disabled="!thrownDice" @click="endTurn">End turn</button>
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
        methods:{
            endTurn(){
                if (!this.thrownDice) {
                    return false;
                }
                let ws = this.$store.getters.ws;
                if (ws === null) {
                    return;
                }
                ws.send(JSON.stringify({
                    type: "end_turn"
                }));
            }
        }
    }
</script>