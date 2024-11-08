/**
 * @brief This class is a wrapper for p5.js so that it can be used to pass to
 * the ResizableP5Canvas custom HTML element. To use this class, simply extend
 * it and implement the setup() and draw() methods. Examlpe can be found in 
 * the visualizer/src/renderer folder.
 */
export default class P5CanvasController {
  constructor(jsonData) {
    this.p5Canvas = null;
    this.p5 = null;
    this.json = jsonData;
    this.unit = 1;
  }

  get width() { return this.p5Canvas.clientWidth; }
  get height() { return this.p5Canvas.clientHeight; }

  updateJson(newJsonData) {this.json = newJsonData;}
  registerCanvas(resizablep5Canvas) {
    this.p5Canvas = resizablep5Canvas.canvas;
    resizablep5Canvas.setController(this);
  }
  hasRegisteredCanvas() {return this.p5Canvas != null;}
  onRedraw() {}
}
