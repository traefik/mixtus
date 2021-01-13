const path = require('path');
const logger = require('./logger');
const fs = require('fs');

/**
 * Function that calculates the conditions to process the file
 * on the iterator based on the include and exclude patterns.
 * @param {string} filePath
 * @param {RegExp} includePattern
 * @param {RegExp} excludePattern
 */
function shouldProcessFile(filePath, includePattern, excludePattern) {
  if (!includePattern && !excludePattern) {
    return true;
  }

  if (
    includePattern && includePattern.test(filePath) &&
    excludePattern && !excludePattern.test(filePath)
  ) {
    return true;
  }

  if (includePattern && !excludePattern && includePattern.test(filePath)) {
    return true;
  }

  if (excludePattern && !includePattern && !excludePattern.test(filePath)) {
    return true;
  }

  return false;
}

/**
 * Function that goes path by path, checks if it's a file
 * or directory, filters based on opts and executes the fileHandler when
 * matching files are found.
 * 
 * @param {string} dirPath
 * @param {function} fileHandler
 * @param {object} opts
 * @param {RegExp} opts.includePattern
 * @param {RegExp} opts.excludePattern
 */
async function iterateDirectory(dirPath, fileHandler, opts) {
  try {
    // get list of files inside dirPath
    const files = await fs.promises.readdir(dirPath);

    // Iterate files
    for (const file of files) {
      const filePath = path.join(dirPath, file);
      const stat = await fs.promises.stat(filePath);

      if (stat.isDirectory()) {
        iterateDirectory(filePath, fileHandler, opts);
      } else if (stat.isFile()) {
        if (shouldProcessFile(filePath, opts.includePattern, opts.excludePattern)) {
          fileHandler(filePath);
        }
      }
    }
  } catch (e) {
    logger.error(e);
  }
}

module.exports = iterateDirectory;