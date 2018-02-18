<template>
    <button v-if="myTurn" type="button" @click="throwDice">Throw the Dice</button>
</template>
<script>
export default {
    computed: {
      myTurn(){
          return this.$store.getters.myTurn;
      }
    },
    methods: {
        throwDice: function(){
            let ws = this.$store.getters.ws;
            if( ws === null){
                return;
            }
            ws.send(JSON.stringify({
                action: "throw_dice"
            }));
        }
    }
}
</script>