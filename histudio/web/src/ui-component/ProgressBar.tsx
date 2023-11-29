import React from 'react';

interface ProgressBarProps {
  value: number;
}

const ProgressBar: React.FC<ProgressBarProps> = ({ value }) => {
  const progressBarStyles = `
    .n-progress {
      width: 100%;
      background-color: #ccc; /* Background color for the progress bar container */
      border-radius: 8px;
    }

    .n-progress-content {
      position: relative;
      height: 20px; /* Height of the progress bar */
    }

    .n-progress-graph {
      position: relative;
      width: 100%; /* Fill the entire container */
      height: 100%;
    }

    .n-progress-graph-line {
      position: absolute;
      width: ${value}%; /* Set the width based on the value */
      height: 100%;
      background: repeating-linear-gradient(
        -45deg,
        #2080f0 3px, /* Use the same color (#2080f0) for both stripes */
        transparent 7px
      );
      border-radius: 8px; /* Add rounded edges */
      background-size: 100% 150%; /* Width should be 2 times of the number of stripes */
      animation: moveStripes 2s linear infinite; /* Adjust the animation duration as needed */
    }

    .n-progress-graph-line-fill {
      width: 100%;
      height: 100%;
    }

    .n-progress-graph-line-indicator {
      position: absolute;
      top: 50%;
      left: ${value}%; /* Position the indicator based on the value */
      transform: translate(-50%, -50%);
      color: #fff; /* Color of the progress indicator text */
      font-size: 14px; /* Font size of the progress indicator text */
    }

    @keyframes moveStripes {
      0% {
        background-position: 0 0;
      }
      100% {
        background-position: 20px 0; /* Adjust the background position to move the stripes */
      }
    }
  `;

  return (
    <div className="n-progress n-progress--line n-progress--default">
      <style>{progressBarStyles}</style>
      <div className="n-progress-content" role="none">
        <div className="n-progress-graph" aria-hidden="true">
          <div className="n-progress-graph-line n-progress-graph-line--indicator-inside">
            <div className="n-progress-graph-line-fill n-progress-graph-line-fill--processing">
              <span className="n-progress-graph-line-indicator">{`${value.toFixed(2)}%`}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default ProgressBar;
