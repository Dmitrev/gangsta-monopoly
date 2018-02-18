<template>
    <p>Current turn {{currentPlayerName}}</p>
</template>
<script>
    export default {
        computed: {
            currentPlayerName(){
                let currentPlayer = this.$store.getters.currentPlayer;

                if( currentPlayer === null ){
                    return null;
                }
                return currentPlayer.name
            }
        }
    }
</script>