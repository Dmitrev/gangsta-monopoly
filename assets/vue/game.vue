<template>
    <canvas id="game"></canvas>
</template>

<script>
    import * as PIXI from 'pixi.js';
    import Game from '../js/game/game'

    export default {
        data(){
          return {
              app: null
          }
        },
        mounted(){
            let canvas = document.getElementById("game");
            new Game(canvas);
        }
    }
</script>