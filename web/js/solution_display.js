function setActiveNav() {
  document.getElementById("homea").classList.remove("is-active");
  document.getElementById("eighta").classList.remove("is-active");
  document.getElementById("twelvea").classList.remove("is-active");
  document.getElementById("sixteena").classList.remove("is-active");
  
  switch(queens) {
    case 8:
      document.getElementById("eighta").classList.add("is-active");
      break;
    case 12:
      document.getElementById("twelvea").classList.add("is-active");
      break;
    case 16:
      document.getElementById("sixteena").classList.add("is-active");
      break;
    default:
      document.getElementById("homea").classList.add("is-active");
  }
}

function drawBoard(canvas, ctx, q, idx) {
  let rect = canvas.getBoundingClientRect();
  let cw = c.width;
  let ch = c.height;
  let sw = cw / q;
  let sh = ch / q;
  let s = solns[idx]["Queens"]

  for (n=0; n<q; n++) {
    let sCol = s[n];
    //console.log(sCol);
    for (m=0; m<q; m++) {
      ctx.beginPath();
      ctx.rect(n*sw, m*sh, sw, sh);
      if (m%2) {
        if (n%2) {
          ctx.fillStyle = "black";
          ctx.fill();
        } else {
          ctx.fillStyle = "white";
          ctx.fill();
        }
      } else {
        if (n%2) {
          ctx.fillStyle = "white";
          ctx.fill();
        } else {
          ctx.fillStyle = "black";
          ctx.fill();
        }
      }
      if (sCol["Row"] == m+1) {
        if (sCol["Moveable"]) {
          ctx.drawImage(goimg, n*sw, m*sh, sw, sh);
        } else {
          ctx.drawImage(stimg, n*sw, m*sh, sw, sh);
        }
      } else if (sCol["Row"] == -1) {
        if (!sCol["Domain"].includes(m+1)) {
          ctx.drawImage(noimg, n*sw, m*sh, sw, sh);
        }
      }
    }
  }
}

function checkButtonDisable() {
  if (idx == 0) {
    document.getElementById('prevButton').disabled = true;
    document.getElementById('nextButton').disabled = false;
  } else if (idx == solns.length-1) {
    document.getElementById('prevButton').disabled = false;
    document.getElementById('nextButton').disabled = true;
  } else {
    document.getElementById('prevButton').disabled = false;
    document.getElementById('nextButton').disabled = false;
  }
}

function updateInfo() {
  let step = document.getElementById('stepInfo');
  let info = document.getElementById('moveInfo');
  step.innerHTML = "Step " + (idx+1);
  info.innerHTML = "Pivot Queen: " + solns[idx]["Column"] + ", Set Value: " + solns[idx]["Value"];
}

function drawPrev() {
  if (idx > 0) {
    idx--;
    drawBoard(c, ctx, queens, idx);
  } 
  checkButtonDisable();
  updateInfo();
}

function drawNext() {
  if (idx < queens) {
    idx++;
    drawBoard(c, ctx, queens, idx);
  } 
  checkButtonDisable();
  updateInfo();
}
