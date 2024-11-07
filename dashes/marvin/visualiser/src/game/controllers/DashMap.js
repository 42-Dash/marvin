import P5CanvasController from './p5Canvas.js';

/**
 * @class This class encapsulate all the location information that a map holds.
 */
class DashMapInformation {
  constructor()
  {
    /** @property {number} squareSize The size of a square in the map in pixel. */
    this.squareSize = 1;
    /** @property {number} squaresDistance The distance between two squares in the map as pixel. */
    this.squaresDistance = 1.1;
    /** @property {number} offsetPercentage The percentage of the square size that is used as offset. */
    this.offsetPercentage = 0.0;
    /** @property {number} leftPadding The left padding of the map in pixel */
    this.leftPadding = 0;
    /** @property {number} UpPadding The up padding of the map in pixel */
    this.UpPadding = 0;
  }

  /**
   * Used only internally inside refresh().
   * @param {number} amount Amount of squares needed
   * @returns The length of the amount of squares in pixel in proportion to the square size.
   */
  #lengthPercentage(amount) {
    if (amount <= 1) return amount;
    return 1 + (1 + this.offsetPercentage) * (amount - 1);
  };

  /**
   * @brief Refreshes the information of the map.
   */
  refresh(mapRowsCount, mapColumnsCount, mapWidth, mapHeight)
  {
    // When the aspect ratio is > 1, it is landscape, otherwise it is portrait.
    const mapAspectRatio = mapColumnsCount / mapRowsCount;
    const screenAspectRatio = mapWidth / mapHeight;
    const mapWidthPercentage = this.#lengthPercentage(mapColumnsCount);
    const mapHeightPercentage = this.#lengthPercentage(mapRowsCount);
    this.leftPadding = 0;
    this.UpPadding = 0;
    if (mapAspectRatio < screenAspectRatio) {
      // Landscape, the height of the map stays the same.
      this.squareSize = mapHeight / mapHeightPercentage;
      this.leftPadding = (mapWidth - mapWidthPercentage * this.squareSize) / 2;
    }
    else {
      // Portrait, the width of the map stays the same.
      this.squareSize = mapWidth / mapWidthPercentage;
      this.UpPadding = (mapHeight - mapHeightPercentage * this.squareSize) / 2;
    }
    this.squaresDistance = this.squareSize * (1 + this.offsetPercentage);
  }
  
  /**
   * Calculates the coordinates of a square given its row and column index.
   * @param {number} i the row index of the square.
   * @param {number} j the column index of the square.
   * @returns An object containing the coordinates of the square.
   */
  squareCoordinates(i, j) {
    const offsetPercentagePlusOne = 1 + this.offsetPercentage;
    return {
      x: j * offsetPercentagePlusOne * this.squareSize + this.leftPadding,
      y: i * offsetPercentagePlusOne * this.squareSize + this.UpPadding
    };
  }

  /**
   * Calculates the coordinates of the center of a square given its row and column index.
   * @param {number} i the row index of the square.
   * @param {number} j the column index of the square.
   * @returns An object containing the coordinates of the center of the square.
   */
  squareCenterCoordinates(i, j) {
    const squareCoordinates = this.squareCoordinates(i, j);
    return {
      x: squareCoordinates.x + this.squareSize / 2,
      y: squareCoordinates.y + this.squareSize / 2
    };
  }
}

/**
 * @class This class is responsible for rendering the map.
 */
export default class DashMapController extends P5CanvasController {
  constructor(jsonMap, league) {
    super(jsonMap);
		this.league = league;
    this.information = new DashMapInformation();
  }
  
  setup() {
    this.p5.createCanvas(this.width, this.height, this.p5Canvas);
    this.p5.frameRate(30);
    this.p5.noLoop();
  }

  draw() {
    this.p5.clear();
    this.p5.stroke([10, 10, 40]);
    this.p5.noStroke();
    
    this.p5.fill(0);
    if (this.league == "Rookie League") {
      this.#drawRookieMap(this.offsetPercentage);
    }
    else {
      this.#drawMap(this.offsetPercentage);
    }
  }

  get mapColumnsCount() {
    if (this.league == "Rookie League") {
      return this.json[0].length;
    }
    else {
      return this.json[0].length / 2;
    }
  }

  get mapRowsCount() { return this.json.length; }
  get startPoint() {
    let mapType = (this.league == "Rookie League") ? 1 : 2;

    for (let i = 0; i < this.mapRowsCount; i++) {
      for (let j = 0; j < this.mapColumnsCount; j++) {
        if (this.json[i][j * mapType] == 'M') {
          return this.information.squareCenterCoordinates(i, j);
        }
      }
    }
    return null
  }

  /**
   * @param {number} value The percentage of the square size that is used as offset.
   */
  set offsetPercentage(value) {
    this.information.offsetPercentage = value;
  }

  #drawRookieMap() {
    this.information.refresh(this.mapRowsCount, this.mapColumnsCount, this.width, this.height);
    // Draw the map
    for (let i = 0; i < this.mapRowsCount; i++) {
      for (let j = 0; j < this.mapColumnsCount; j++) {
        const value = this.json[i][j];
        if (value > 0 && value < 10) {
          const weight = 255 - value * 255/10 + 20;
          this.p5.fill(weight);
        } 
        else if (value == 'G') {
          this.p5.fill(255, 110, 100);
        }
        else if (value == 'M') {
          this.p5.fill(100, 110, 255);
          
        }
        const {x, y} = this.information.squareCoordinates(i, j);
        this.p5.rect(x, y, this.information.squareSize, this.information.squareSize);
      }
    }
  }


	#calcTileColour(x, y) {
		let hue = this.json[y][x * 2];
		let luminance = parseInt(this.json[y][(x * 2) + 1]);
	
		// console.log("hue: ", hue, " luminance: ", luminance);
		let color;
		switch (hue) {
			case 'W':
				color = [10, 153, 255];
				break;
			case 'A':
				color = [250, 250, 220];
				break;
			case 'E':
				color = [100, 200, 100];
				break;
			case 'M':
				color = [255, 200, 0];
				break;
			case 'G':
				color = [0, 0, 0];
				break;
			default:
				color = [100, 100, 100];
		}
	
		const luminanceValue = 1 - luminance / 12;
		if (!isNaN(luminanceValue)) {
			color = [
				color[0] * luminanceValue,
				color[1] * luminanceValue,
				color[2] * luminanceValue
			];
		}
		return color;
	}

	#drawMap() {
		this.information.refresh(this.mapRowsCount, this.mapColumnsCount, this.width, this.height);

		// Draw the map
		for (let i = 0; i < this.mapRowsCount; i++) {
			for (let j = 0; j < this.mapColumnsCount; j++) {
				// Correct the way to call the calc method using 'this'
				this.tileColour = this.#calcTileColour(j, i);

				this.p5.fill(this.tileColour); // Access tileColour using this

				const { x, y } = this.information.squareCoordinates(i, j);
				this.p5.rect(x, y, this.information.squareSize, this.information.squareSize);
			}
		}
	}
}
