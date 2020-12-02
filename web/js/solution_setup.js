function drawBoard(canvas, ctx, q) {
  let cw = c.width;
  let ch = c.height;
  let sw = cw / q;
  let sh = ch / q;

  for (n=0; n<q; n++) {
    for (m=0; m<q; m++) {
      ctx.beginPath();
      ctx.rect(n*sw, m*sh, sw, sh);
      if (n%2) {
        if (m%2) {
          ctx.fillStyle = "black";
          ctx.fill();
        } else {
          ctx.fillStyle = "white";
          ctx.fill();
        }
      } else {
        if (m%2) {
          ctx.fillStyle = "white";
          ctx.fill();
        } else {
          ctx.fillStyle = "black";
          ctx.fill();
        }
      }
    }
  }
}

function drawCol(bm, bn, q) {
  let cw = c.width;
  let ch = c.height;
  let sw = cw / q;
  let sh = ch / q;

  for (m=0; m<q; m++) {
    ctx.beginPath();
    ctx.rect(bn*sw, m*sh, sw, sh);
    if (bn%2) {
      if (m%2) {
        ctx.fillStyle = "black";
        ctx.fill();
      } else {
        ctx.fillStyle = "white";
        ctx.fill();
      }
    } else {
      if (m%2) {
        ctx.fillStyle = "white";
        ctx.fill();
      } else {
        ctx.fillStyle = "black";
        ctx.fill();
      }
    }
    if (m==bm-1) {
      ctx.drawImage(stimg, bn*sw, m*sh, sw, sh);
    } 
  }
}