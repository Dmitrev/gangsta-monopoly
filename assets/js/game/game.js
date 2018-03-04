class Game {
    constructor(canvas){
        this.canvas = canvas;
        this.createApplication();
        this.loadFiles();
    }

    createApplication(){
        this.app = new PIXI.Application({
            width: this.canvas.width,
            height: this.canvas.height,
            view: this.canvas,
            backgroundColor: 0xff00ff
        });

        this.app.renderer.view.style.position = "absolute";
        this.app.renderer.view.style.display = "block";
        this.app.renderer.autoResize = true;
        this.app.renderer.resize(window.innerWidth, window.innerHeight);
    }

    loadFiles(){
        PIXI.loader
            .add('../img/cat.png')
            .load(() => this.setup());
    }
    setup(){
        let sprite = new PIXI.Sprite(
            PIXI.loader.resources["../img/cat.png"].texture
        );

        this.app.stage.addChild(sprite);
    }
}

export default Game;