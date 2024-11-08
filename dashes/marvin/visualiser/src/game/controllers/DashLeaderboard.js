export default class DashLeaderboard {
  constructor(gameData, leaderboard) {
    this.gameData = gameData;
    this.leaderboard = leaderboard;
    this.started = false;
  }

  compareGroups(groupB, groupA) {
    const statusOrder = { "valid": 0, "timeout": 1, "invalid": 2 };
    if (statusOrder[groupA.status] > statusOrder[groupB.status]) {
      return -1;
    }
    else if (statusOrder[groupA.status] < statusOrder[groupB.status]) {
      return 1;
    }
    else if (groupA.score > groupB.score) {
      return -1;
    }
    else if (groupA.score < groupB.score) {
      return 1;
    }
    return groupB.name.localeCompare(groupA.name);
  }

  showCurrentPoints() {
    this.leaderboard.showCurrentPoints();
  }

  hideCurrentPoints() {
    this.leaderboard.hideCurrentPoints();
  }

  resetCost() {
    this.leaderboard.resetCost();
  }

  renderDefaultLeaderboard() {
    this.leaderboard.hideCurrentPoints();
    this.leaderboard.loadRanking(
    this.gameData.groups.map((group, index) => {
      return {
        name: group.name,
        total_score: 0,
        current_points: 0,
        cost: 0,
        status: "valid",
        colour: this.gameData.colorByGroupName(group.name),
      };
    }));
  }

  renderLeaderboard() {
    const ranking = new Map();

    this.gameData.groups.forEach((group, index) => {
      ranking.set(group.name, {
        total_score: 0,
        current_points: 0,
        cost: 0,
        status: group.status,
        colour: this.gameData.colorByGroupName(group.name),
      });
    });

    for (let level = 0; this.started && level <= this.gameData.level; level++) {
      const levelGroups = this.gameData.groupsAt(level);
      levelGroups.sort(this.compareGroups);

      // Compute points for this level
      for (let i = 0, points = levelGroups.length + 1; i < levelGroups.length; i++) {
        const group = levelGroups[i];
        const savedGroupInfo = ranking.get(group.name);

        if (group.status != 'valid') {
          ranking.set(group.name, {
              ...savedGroupInfo,
              status: group.status,
              cost: group.score,
              current_points: 0,
            });
            break;
        }

        const previousGroupInfo = i > 0 ? ranking.get(levelGroups[i - 1].name) : null;

        if (previousGroupInfo != null
          && (group.score != previousGroupInfo.cost || previousGroupInfo.status != "valid")) {
          points--;
        }

        ranking.set(group.name, {
          ...savedGroupInfo,
          status: group.status,
          cost: group.score,
          current_points: points,
          total_score: savedGroupInfo.total_score + points
        });
      }
    }
    this.started = true;

    const rankedGroups = [];
    ranking.forEach((value, key, map) => {
      rankedGroups.push({
        name: key,
        ...value
      });
    });
    rankedGroups.sort((a, b) => b.total_score - a.total_score);
    this.leaderboard.loadRanking(rankedGroups);
    this.leaderboard.showCurrentPoints();
  }
}
