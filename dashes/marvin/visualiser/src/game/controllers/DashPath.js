import P5CanvasController from './p5Canvas.js';

export default class DashPathController extends P5CanvasController {
  constructor(jsonPath, dashMap, color) {
    super(jsonPath)
    this.dashMap = dashMap;
    this.color = color;
	
    this.lineSpeed = 10;
    this.stepcount = 0;
    this.lines = [];
    this.lineLength = 0;
    this.started = false;
  }

  setup() {
    this.p5.createCanvas(this.width, this.height, this.p5Canvas);
    this.p5.frameRate(30);

    this.lines.push({
      x: this.startPoint.x,
      y: this.startPoint.y
    })
    this.p5.noLoop();
	}

	calcStrokeWeight() {
		let strokeWeight = this.dashMap.information.squareSize / 10;
		if (strokeWeight < 5) {
			strokeWeight = 5;
		}
		return strokeWeight;
	}

  draw() {
    if (!this.started) return;

    this.p5.clear();
    this.p5.stroke(this.color.r, this.color.g, this.color.b);
    this.p5.strokeWeight(this.calcStrokeWeight());
    
    if (this.stepcount >= this.json.length) {
      this.started = false;
      this.p5.noLoop();
    }
    
    // print all done line first
    for (let i = 0; i < this.lines.length - 1; i++) {
      this.p5.line(this.lines[i].x, this.lines[i].y, this.lines[i + 1].x, this.lines[i + 1].y);
    }
    
    let lineDirection = this.lineDirection;
    if (lineDirection == null) {
      this.started = false;
      this.p5.noLoop();
      return ;
    }
    let curPoint = this.lines[this.stepcount];
    
    this.x2 = curPoint.x + this.lineDirection.x * this.lineLength;
    this.y2 = curPoint.y + this.lineDirection.y * this.lineLength;
    
		this.p5.circle(this.x2, this.y2, this.calcStrokeWeight())
    this.p5.line(curPoint.x, curPoint.y, this.x2, this.y2);
    
    this.lineLength += this.lineSpeed;
    
    // direction change
    if (this.lineLength > this.squaresDistance) {
      this.lineLength = 0;
      this.stepcount++;
      this.lines[this.stepcount] = {
        x: curPoint.x + this.squaresDistance * lineDirection.x,
        y: curPoint.y + this.squaresDistance * lineDirection.y
      };
    }
  }
  
  onRedraw() {
    this.#reset();
    this.p5.loop();
  }

  get startPoint() { return this.dashMap.startPoint; }
  get squaresDistance() { return this.dashMap.information.squaresDistance; }
  get lineDirection() {
    const directionChar = this.json[this.stepcount];

    switch (directionChar) {
      case 'U':
      return {x: 0, y: -1}; // Up
      case 'D':
      return {x: 0, y: 1};  // Down
      case 'L':
      return {x: -1, y: 0}; // Left
      case 'R':
      return {x: 1, y: 0}  // Right
      default:
      return null  // No movement for an unknown direction
    }
  }

  start() {
    this.started = true;
    this.p5.loop();
  }
  
  clear() {
    this.started = false;
    this.#reset();
    this.p5.noLoop();
  }

  #reset() {
    this.p5.clear();
    this.lines = [];
    this.stepcount = 0;
    this.lineLength = 0;
    this.lines.push({
      x: this.startPoint.x,
      y: this.startPoint.y
    });
  }
}
