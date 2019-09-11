const {FieldAlignment, FieldDynamicContentAlignment} = require('./types');

function alignOffsetToType(off, fieldType) {
    const fieldSize = FieldAlignment[fieldType];
    return Math.floor((off + fieldSize - 1) / fieldSize) * fieldSize;
}

function alignDynamicFieldContentOffset(off, fieldType) {
    const contentAlignment = FieldDynamicContentAlignment[fieldType];
    return Math.floor((off + contentAlignment - 1) / contentAlignment) * contentAlignment;
}

module.exports = {alignDynamicFieldContentOffset, alignOffsetToType};
