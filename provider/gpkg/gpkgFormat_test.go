package gpkg

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/terranodo/tegola"
	"fmt"
)

var geometryBlob = []byte{71,80,0,3,230,16,0,0,229,109,250,182,103,182,55,64,193,171,176,208,185,203,55,64,44,201,188,229,214,242,66,64,32,194,46,134,184,248,66,64,1,3,0,0,0,1,0,0,0,15,1,0,0,100,160,98,106,172,183,55,64,44,201,188,229,214,242,66,64,80,98,107,17,172,183,55,64,20,84,249,67,216,242,66,64,74,157,54,153,167,183,55,64,18,99,170,170,225,242,66,64,109,146,122,170,158,183,55,64,61,197,187,109,240,242,66,64,111,230,161,0,155,183,55,64,117,19,124,211,244,242,66,64,67,234,118,246,149,183,55,64,107,226,63,56,250,242,66,64,125,231,6,249,143,183,55,64,120,34,45,58,254,242,66,64,36,36,119,125,139,183,55,64,24,250,22,140,0,243,66,64,8,227,241,254,137,183,55,64,216,55,82,17,2,243,66,64,66,205,144,42,138,183,55,64,169,9,162,238,3,243,66,64,68,1,219,193,136,183,55,64,89,30,107,70,6,243,66,64,229,188,90,73,134,183,55,64,35,117,2,154,8,243,66,64,99,130,26,190,133,183,55,64,241,136,209,115,11,243,66,64,203,136,193,178,136,183,55,64,21,63,215,8,14,243,66,64,131,125,224,207,149,183,55,64,82,252,223,17,21,243,66,64,163,79,80,47,157,183,55,64,162,3,129,132,23,243,66,64,139,107,107,250,162,183,55,64,0,32,39,76,24,243,66,64,127,97,141,14,163,183,55,64,23,0,153,140,25,243,66,64,93,52,117,104,162,183,55,64,242,245,107,161,26,243,66,64,68,19,205,168,158,183,55,64,39,239,136,175,27,243,66,64,238,184,77,93,158,183,55,64,162,246,199,32,28,243,66,64,239,16,255,176,165,183,55,64,252,223,108,206,28,243,66,64,186,116,130,236,171,183,55,64,191,147,113,231,29,243,66,64,51,199,185,77,184,183,55,64,8,196,53,136,32,243,66,64,4,51,58,123,194,183,55,64,28,224,238,81,36,243,66,64,24,251,203,238,201,183,55,64,187,157,51,88,39,243,66,64,245,234,153,20,213,183,55,64,31,127,14,152,44,243,66,64,51,129,17,234,216,183,55,64,218,151,125,104,48,243,66,64,61,107,18,23,219,183,55,64,14,139,98,3,51,243,66,64,141,33,182,170,219,183,55,64,230,44,14,194,55,243,66,64,153,239,76,13,217,183,55,64,161,84,251,116,60,243,66,64,10,210,140,69,211,183,55,64,201,192,183,159,66,243,66,64,146,108,100,13,205,183,55,64,81,172,191,202,70,243,66,64,96,27,150,171,196,183,55,64,210,34,71,223,73,243,66,64,61,217,148,43,188,183,55,64,251,57,96,178,76,243,66,64,151,147,171,179,181,183,55,64,113,146,156,2,79,243,66,64,113,25,237,167,181,183,55,64,188,225,244,228,80,243,66,64,58,159,223,156,182,183,55,64,225,177,159,197,82,243,66,64,39,209,72,58,185,183,55,64,110,50,5,198,85,243,66,64,192,39,66,143,189,183,55,64,72,146,49,136,89,243,66,64,177,208,189,61,190,183,55,64,39,221,167,244,93,243,66,64,253,89,152,42,189,183,55,64,86,105,65,222,97,243,66,64,172,111,170,255,189,183,55,64,162,0,81,48,99,243,66,64,211,226,231,26,193,183,55,64,234,152,169,255,98,243,66,64,242,159,110,160,192,183,55,64,8,164,31,104,96,243,66,64,216,67,251,88,193,183,55,64,112,163,18,232,93,243,66,64,31,116,191,249,195,183,55,64,186,249,235,186,92,243,66,64,39,42,118,143,199,183,55,64,175,231,16,221,88,243,66,64,73,68,157,94,206,183,55,64,101,125,202,49,89,243,66,64,128,198,167,91,209,183,55,64,29,249,222,58,90,243,66,64,37,38,3,155,206,183,55,64,10,216,105,255,94,243,66,64,129,190,170,105,205,183,55,64,229,15,188,16,97,243,66,64,194,60,43,105,197,183,55,64,44,76,240,151,96,243,66,64,224,249,177,238,196,183,55,64,193,149,253,4,97,243,66,64,66,134,233,214,198,183,55,64,87,151,83,2,98,243,66,64,160,131,46,225,208,183,55,64,196,76,253,23,99,243,66,64,246,39,93,159,222,183,55,64,166,65,135,175,101,243,66,64,57,247,13,2,225,183,55,64,207,186,252,61,103,243,66,64,132,144,226,210,231,183,55,64,136,54,17,71,104,243,66,64,94,122,69,1,236,183,55,64,103,115,82,78,106,243,66,64,58,92,171,61,236,183,55,64,75,176,147,85,108,243,66,64,172,3,32,238,234,183,55,64,79,103,147,104,110,243,66,64,37,77,222,91,236,183,55,64,79,103,147,104,110,243,66,64,171,31,138,184,239,183,55,64,229,254,143,184,108,243,66,64,211,126,90,251,241,183,55,64,182,82,25,16,107,243,66,64,66,235,3,36,245,183,55,64,80,233,204,226,106,243,66,64,91,7,189,237,248,183,55,64,27,96,65,245,106,243,66,64,110,207,78,97,0,184,55,64,69,13,1,27,107,243,66,64,84,195,143,123,4,184,55,64,182,56,116,196,107,243,66,64,80,158,195,198,6,184,55,64,163,179,221,78,108,243,66,64,159,32,29,195,8,184,55,64,93,165,187,235,108,243,66,64,19,116,106,29,11,184,55,64,219,100,67,237,109,243,66,64,244,103,171,55,15,184,55,64,0,196,19,48,112,243,66,64,71,182,186,156,18,184,55,64,41,81,246,150,114,243,66,64,55,135,16,252,20,184,55,64,88,89,145,135,116,243,66,64,147,158,199,205,21,184,55,64,239,187,250,194,117,243,66,64,160,121,74,24,23,184,55,64,175,54,198,78,120,243,66,64,216,114,103,38,24,184,55,64,144,181,134,82,123,243,66,64,216,47,159,172,24,184,55,64,245,116,191,84,125,243,66,64,59,220,179,83,24,184,55,64,88,20,27,152,129,243,66,64,237,42,255,181,23,184,55,64,7,15,63,164,132,243,66,64,54,83,198,100,22,184,55,64,130,128,215,194,135,243,66,64,139,225,143,71,21,184,55,64,105,122,80,80,138,243,66,64,249,71,206,120,17,184,55,64,58,175,120,234,145,243,66,64,137,152,92,214,14,184,55,64,196,226,55,133,149,243,66,64,151,194,23,129,12,184,55,64,75,186,210,215,152,243,66,64,139,196,169,123,9,184,55,64,8,211,65,168,156,243,66,64,23,73,130,112,5,184,55,64,107,114,157,235,160,243,66,64,205,38,192,176,252,183,55,64,19,178,78,240,168,243,66,64,53,211,189,78,234,183,55,64,138,7,148,77,185,243,66,64,129,32,81,178,230,183,55,64,136,205,17,67,188,243,66,64,68,225,14,47,227,183,55,64,242,27,112,167,190,243,66,64,0,241,4,12,220,183,55,64,108,135,208,82,195,243,66,64,139,234,247,14,200,183,55,64,5,177,51,133,206,243,66,64,47,210,196,59,192,183,55,64,115,70,0,220,209,243,66,64,194,51,178,117,186,183,55,64,189,52,69,128,211,243,66,64,206,120,91,233,181,183,55,64,87,105,156,151,212,243,66,64,145,17,63,181,176,183,55,64,198,240,51,137,213,243,66,64,227,223,194,21,171,183,55,64,19,254,12,111,214,243,66,64,81,42,151,124,162,183,55,64,52,101,167,31,212,243,66,64,27,168,140,127,159,183,55,64,27,101,88,32,213,243,66,64,42,215,54,32,157,183,55,64,111,125,88,111,212,243,66,64,190,211,157,39,158,183,55,64,188,170,14,20,211,243,66,64,188,98,144,5,150,183,55,64,230,47,244,136,209,243,66,64,58,72,45,57,147,183,55,64,165,18,84,67,209,243,66,64,182,229,18,253,144,183,55,64,64,123,245,241,208,243,66,64,154,72,105,54,143,183,55,64,46,134,205,91,208,243,66,64,57,240,123,229,139,183,55,64,206,151,57,184,207,243,66,64,238,134,126,183,138,183,55,64,120,220,166,46,207,243,66,64,240,186,200,78,137,183,55,64,148,245,81,252,206,243,66,64,199,19,65,156,135,183,55,64,95,108,198,14,207,243,66,64,54,127,110,195,131,183,55,64,66,155,210,176,206,243,66,64,100,142,138,164,130,183,55,64,107,144,73,70,206,243,66,64,228,68,204,54,129,183,55,64,7,249,234,244,205,243,66,64,38,79,72,33,127,183,55,64,207,19,59,191,205,243,66,64,98,42,105,106,126,183,55,64,78,171,232,15,205,243,66,64,94,181,232,83,125,183,55,64,78,243,159,127,204,243,66,64,107,243,16,215,123,183,55,64,137,37,246,26,204,243,66,64,137,183,24,77,120,183,55,64,163,221,141,170,203,243,66,64,222,253,42,192,119,183,55,64,247,153,105,177,202,243,66,64,95,205,200,32,119,183,55,64,215,200,117,83,202,243,66,64,228,59,83,67,118,183,55,64,233,77,12,201,201,243,66,64,17,75,111,36,117,183,55,64,199,211,77,189,201,243,66,64,106,242,148,213,116,183,55,64,43,107,172,14,202,243,66,64,80,130,180,181,116,183,55,64,155,150,31,184,202,243,66,64,245,193,50,54,116,183,55,64,92,49,200,2,203,243,66,64,123,33,63,118,114,183,55,64,50,178,26,1,203,243,66,64,19,47,5,90,112,183,55,64,8,5,91,219,202,243,66,64,30,96,65,245,106,183,55,64,133,38,63,152,202,243,66,64,174,109,7,217,104,183,55,64,185,175,202,133,202,243,66,64,253,42,44,93,102,183,55,64,104,85,75,58,202,243,66,64,71,155,170,123,100,183,55,64,9,241,237,2,202,243,66,64,203,205,237,20,97,183,55,64,4,190,236,232,201,243,66,64,147,3,44,168,94,183,55,64,102,157,2,170,201,243,66,64,196,55,20,62,91,183,55,64,235,149,195,56,201,243,66,64,16,220,220,243,87,183,55,64,146,218,48,175,200,243,66,64,194,21,63,124,79,183,55,64,154,157,160,168,199,243,66,64,94,84,65,251,71,183,55,64,237,43,106,139,198,243,66,64,76,212,102,247,63,183,55,64,40,166,119,150,197,243,66,64,175,48,199,60,60,183,55,64,183,122,4,237,196,243,66,64,3,44,20,68,56,183,55,64,94,191,113,99,196,243,66,64,125,176,157,57,53,183,55,64,249,39,19,18,196,243,66,64,164,238,20,188,50,183,55,64,224,255,233,97,195,243,66,64,164,250,132,162,47,183,55,64,109,43,172,10,195,243,66,64,252,108,228,186,41,183,55,64,24,25,228,46,194,243,66,64,207,152,147,97,38,183,55,64,94,39,6,146,193,243,66,64,123,94,241,212,35,183,55,64,37,133,30,214,192,243,66,64,94,172,94,52,26,183,55,64,107,219,247,168,191,243,66,64,210,120,159,153,22,183,55,64,37,93,68,37,191,243,66,64,119,196,141,0,19,183,55,64,139,132,17,86,190,243,66,64,78,56,244,22,15,183,55,64,213,146,51,185,189,243,66,64,110,237,125,170,10,183,55,64,230,115,238,118,189,243,66,64,87,249,158,145,8,183,55,64,156,9,168,203,189,243,66,64,84,137,13,113,7,183,55,64,217,79,107,63,190,243,66,64,66,246,65,150,5,183,55,64,228,63,164,223,190,243,66,64,109,149,204,86,3,183,55,64,67,210,19,59,191,243,66,64,163,62,53,3,1,183,55,64,156,49,130,124,191,243,66,64,168,139,219,104,0,183,55,64,129,70,233,210,191,243,66,64,74,131,162,121,0,183,55,64,251,77,40,68,192,243,66,64,23,86,59,212,0,183,55,64,84,219,168,169,192,243,66,64,221,107,156,168,0,183,55,64,216,185,196,236,192,243,66,64,6,215,220,209,255,182,55,64,238,41,165,12,193,243,66,64,110,208,151,222,254,182,55,64,115,8,193,79,193,243,66,64,63,153,161,76,254,182,55,64,95,85,24,182,193,243,66,64,15,93,188,196,253,182,55,64,11,15,6,67,194,243,66,64,122,178,155,25,253,182,55,64,72,131,219,218,194,243,66,64,56,11,197,103,252,182,55,64,227,209,215,61,195,243,66,64,212,74,67,232,251,182,55,64,226,45,252,133,195,243,66,64,193,218,98,200,251,182,55,64,11,55,224,243,195,243,66,64,33,63,192,255,251,182,55,64,71,125,163,103,196,243,66,64,206,161,120,58,252,182,55,64,94,211,222,59,197,243,66,64,189,141,188,98,252,182,55,64,181,70,186,85,198,243,66,64,175,12,204,101,254,182,55,64,224,145,29,192,199,243,66,64,59,84,248,216,2,183,55,64,81,71,199,213,200,243,66,64,102,40,73,50,6,183,55,64,252,92,217,170,201,243,66,64,117,143,199,103,13,183,55,64,122,212,169,60,203,243,66,64,190,192,7,10,23,183,55,64,226,218,80,49,206,243,66,64,124,85,120,225,24,183,55,64,209,39,168,151,206,243,66,64,40,55,64,31,27,183,55,64,167,50,49,2,207,243,66,64,92,160,238,77,29,183,55,64,13,130,216,227,207,243,66,64,159,44,215,54,32,183,55,64,122,173,75,141,208,243,66,64,223,199,61,2,36,183,55,64,165,18,84,67,209,243,66,64,164,13,118,121,41,183,55,64,190,58,125,243,209,243,66,64,217,131,16,144,47,183,55,64,83,152,247,56,211,243,66,64,251,97,240,213,51,183,55,64,74,49,172,135,212,243,66,64,1,235,221,196,53,183,55,64,227,219,204,50,213,243,66,64,44,171,193,69,56,183,55,64,8,152,10,59,214,243,66,64,38,109,153,194,57,183,55,64,218,157,164,175,214,243,66,64,115,167,59,79,60,183,55,64,176,168,45,26,215,243,66,64,120,115,241,183,61,183,55,64,181,9,65,88,215,243,66,64,197,225,221,219,62,183,55,64,187,152,102,186,215,243,66,64,154,113,174,188,63,183,55,64,145,117,221,0,216,243,66,64,182,151,69,114,67,183,55,64,245,150,114,190,216,243,66,64,174,77,173,8,72,183,55,64,142,39,238,29,218,243,66,64,86,119,44,182,73,183,55,64,70,25,204,186,218,243,66,64,95,144,136,132,74,183,55,64,78,50,40,137,219,243,66,64,191,105,102,210,75,183,55,64,102,90,81,57,220,243,66,64,157,187,93,47,77,183,55,64,118,79,121,207,220,243,66,64,201,83,103,255,77,183,55,64,140,119,162,127,221,243,66,64,170,190,186,42,80,183,55,64,115,119,83,128,222,243,66,64,134,200,250,23,82,183,55,64,198,137,27,92,223,243,66,64,214,233,64,214,83,183,55,64,43,33,122,173,223,243,66,64,228,196,195,32,85,183,55,64,46,125,158,245,223,243,66,64,149,54,250,61,86,183,55,64,40,2,230,107,224,243,66,64,220,134,155,157,86,183,55,64,223,197,177,228,224,243,66,64,192,206,3,14,87,183,55,64,233,135,216,96,225,243,66,64,163,89,52,248,86,183,55,64,145,14,197,211,225,243,66,64,35,123,24,181,86,183,55,64,97,20,95,72,226,243,66,64,25,118,41,191,86,183,55,64,168,146,18,204,226,243,66,64,93,121,36,185,87,183,55,64,46,159,64,51,227,243,66,64,189,181,167,63,86,183,55,64,231,117,48,7,230,243,66,64,219,74,84,20,84,183,55,64,111,169,239,161,233,243,66,64,242,34,201,98,81,183,55,64,222,82,41,209,237,243,66,64,70,25,39,116,77,183,55,64,89,118,210,12,243,243,66,64,183,239,246,197,74,183,55,64,200,175,122,27,246,243,66,64,170,149,100,120,71,183,55,64,186,245,80,145,249,243,66,64,40,170,92,77,67,183,55,64,114,219,190,71,253,243,66,64,236,91,156,231,62,183,55,64,237,4,160,246,0,244,66,64,76,124,181,163,56,183,55,64,159,203,121,107,5,244,66,64,6,155,41,99,50,183,55,64,149,252,153,139,9,244,66,64,64,24,46,97,35,183,55,64,151,57,93,22,19,244,66,64,157,219,166,211,21,183,55,64,182,31,58,78,27,244,66,64,193,188,125,12,12,183,55,64,119,147,132,233,32,244,66,64,83,104,180,176,2,183,55,64,201,219,91,219,37,244,66,64,170,210,22,215,248,182,55,64,132,54,74,168,42,244,66,64,123,134,55,107,240,182,55,64,237,126,112,153,46,244,66,64,253,210,236,203,227,182,55,64,139,28,216,224,51,244,66,64,158,61,60,240,214,182,55,64,65,68,197,147,56,244,66,64,47,125,84,104,197,182,55,64,229,250,136,169,62,244,66,64,22,8,110,73,179,182,55,64,214,78,148,132,68,244,66,64,196,28,4,29,173,182,55,64,99,49,86,64,70,244,66,64,56,228,85,140,169,182,55,64,139,196,112,26,71,244,66,64,20,230,152,135,167,182,55,64,48,24,92,115,71,244,66,64,97,158,206,21,165,182,55,64,27,55,161,181,71,244,66,64,38,140,85,57,163,182,55,64,103,48,13,195,71,244,66,64,52,182,16,228,160,182,55,64,99,253,11,169,71,244,66,64,140,53,92,228,158,182,55,64,163,98,99,94,71,244,66,64,250,128,172,76,157,182,55,64,0,220,118,235,70,244,66,64,21,201,178,187,155,182,55,64,41,117,201,56,70,244,66,64,95,10,214,56,155,182,55,64,148,43,188,203,69,244,66,64,179,167,29,254,154,182,55,64,109,39,199,83,69,244,66,64,238,125,79,81,154,182,55,64,227,226,168,220,68,244,66,64,169,122,84,87,153,182,55,64,232,93,97,102,68,244,66,64,193,122,163,86,152,182,55,64,64,215,116,243,67,244,66,64,132,119,168,92,151,182,55,64,40,221,93,103,67,244,66,64,161,216,10,154,150,182,55,64,121,76,147,8,67,244,66,64,130,243,169,99,149,182,55,64,124,245,93,182,66,244,66,64,199,22,130,28,148,182,55,64,96,174,160,196,66,244,66,64,5,58,90,213,146,182,55,64,58,190,24,37,67,244,66,64,239,109,85,109,146,182,55,64,52,67,96,155,67,244,66,64,39,83,5,163,146,182,55,64,101,127,69,35,68,244,66,64,222,22,209,27,147,182,55,64,231,185,133,174,68,244,66,64,149,146,229,36,148,182,55,64,243,215,208,114,69,244,66,64,15,4,126,67,151,182,55,64,134,73,184,144,71,244,66,64,229,220,129,149,158,182,55,64,92,230,116,89,76,244,66,64,112,125,196,84,159,182,55,64,250,103,114,214,76,244,66,64,26,62,51,210,156,182,55,64,162,232,38,214,78,244,66,64,4,74,84,185,154,182,55,64,177,31,206,104,80,244,66,64,92,12,104,51,152,182,55,64,194,40,99,215,81,244,66,64,73,81,194,167,148,182,55,64,147,158,142,108,83,244,66,64,90,175,199,233,144,182,55,64,205,38,209,220,84,244,66,64,162,140,201,44,140,182,55,64,38,82,245,134,86,244,66,64,66,59,93,204,133,182,55,64,228,230,101,94,88,244,66,64,159,66,26,186,126,182,55,64,47,54,190,64,90,244,66,64,163,102,106,109,117,182,55,64,99,159,108,111,92,244,66,64,233,49,145,210,108,182,55,64,125,9,21,28,94,244,66,64,229,109,250,182,103,182,55,64,11,174,110,248,94,244,66,64,229,109,250,182,103,182,55,64,32,194,46,134,184,248,66,64,193,171,176,208,185,203,55,64,32,194,46,134,184,248,66,64,193,171,176,208,185,203,55,64,44,201,188,229,214,242,66,64,100,160,98,106,172,183,55,64,44,201,188,229,214,242,66,64}
var geometryHeader = geometryBlob[:40]
var geometryData = geometryBlob[40:]
var wkbLineStringBlobs = [][]byte{
	{1,2,0,0,0,2,0,0,0,212,238,252,71,75,180,55,64,69,93,16,54,225,244,66,64,154,247,238,69,106,180,55,64,69,122,246,1,237,244,66,64},
	{1,2,0,0,0,2,0,0,0,84,166,61,202,45,188,55,64,215,64,169,172,92,245,66,64,16,148,219,246,61,188,55,64,229,26,176,245,86,245,66,64},
}

func TestBytesToFloat64(t *testing.T) {
	// bytes and the float64 value they should be converted to
	testCases := map[[8]byte]float64{
		[8]byte{0,0,0,0,0,0,0,0}: 0,
		[8]byte{64,64,62,86,4,24,147,117}: 32.487,
		[8]byte{64,84,159,43,2,12,73,186}: 82.487,
		[8]byte{63,240,0,0,0,0,0,0}: 1,
		[8]byte{64,248,106,8,0,0,0,0}: 100000.5,
	}
	
	for bytes, expectedValue := range testCases {
		f64Value := bytesToFloat64(bytes[:], wkbXDR)
		if f64Value != expectedValue {
			assert.Equal(t, f64Value, expectedValue, "")
		}
	}
}

func TestBytesToInt32(t *testing.T) {
	testCases := map[[4]byte]int32{
		[4]byte{0,0,0,0}: 0,
		[4]byte{230,16,0,0}: 4326,
	}
	
	for bytes, expectedValue := range testCases {
		i32value := bytesToInt32(bytes[:], wkbNDR)
		assert.Equal(t, i32value, expectedValue, "")
	}
}


func TestGeoPackageBinaryHeaderInit(t *testing.T) {
	initializedExpected := true
	var magicExpected uint16 = 0x4750
	var versionExpected uint8 = 0x0
	var flagsExpected uint8 = 0x3
	var srs_idExpected int32 = 4326
	envelopeExpected := []float64{23.712520061626396, 23.79580406487708, 37.89718314855631, 37.94313123019333}
	sizeExpected := 40

	var h GeoPackageBinaryHeader
	h.Init(geometryHeader)
	
	assert.Equal(t, initializedExpected, h.isInitialized("TestInit"), "")
	assert.Equal(t, magicExpected, h.Magic(), "")
	assert.Equal(t, versionExpected, h.Version(), "")
	assert.Equal(t, flagsExpected, h.flags, "")
	assert.Equal(t, srs_idExpected, h.SRSId(), "")
	assert.Equal(t, envelopeExpected, h.Envelope(), "")
	assert.Equal(t, sizeExpected, h.headerSize)
}

func TestEnvelopeType(t *testing.T) {
	/* The envelope is a 3-bit unsiged integer composed of the flag bits 1-3
	*/
	var h GeoPackageBinaryHeader
	h.flagsReady = true
	
	// The first byte here shouldn't make any difference
	testCases := map[byte]uint8 {
		0xFF: 7,
		0xFE: 7,
		0x3D: 6,
		0x8C: 6,
		0x3B: 5,
		0x2A: 5,
		0x19: 4,
		0x18: 4,
		0x77: 3,
		0x76: 3,
		0x65: 2,
		0x64: 2,
		0x53: 1,
		0x52: 1,
	}

	for flags, expectedEnvType := range testCases {
		h.flags = flags
		envType := h.EnvelopeType()
		assert.Equal(t, envType, expectedEnvType, "")
	}
}

func TestReadGeometries(t *testing.T) {
	geoms, bytesConsumed := readGeometries(geometryData)
	assert.Equal(t, bytesConsumed, len(geometryData), "")
	assert.Equal(t, len(geoms), 1, "")
	// 3 is the expected type for a Polygon
	polygonType := uint32(3)
	assert.Equal(t, polygonType, geoms[0].Type(), "")
}

func TestReadNextGeometry(t *testing.T) {
	g, bytesConsumed := readNextGeometry(geometryData)
	assert.Equal(t, g.Type(), uint32(3), "")
	assert.Equal(t, bytesConsumed, len(geometryData), "")
}


func TestWKBLineString(t *testing.T) {
	expectedSubpoints := [][]tegola.Point{
		{WKBPoint{x:23.7042737, y:37.9131229}, WKBPoint{x:23.7047466, y:37.9134829}},
		{WKBPoint{x:23.7350737, y:37.9168907}, WKBPoint{x:23.7353205, y:37.9167163}},
	}

	expectedTypes := []uint32{2, 2}
	for i := 0; i < len(wkbLineStringBlobs); i++ {
		ls := new(WKBLineString)
		bytesConsumed := ls.Init(wkbLineStringBlobs[i])
		assert.Equal(t, len(wkbLineStringBlobs[i]), bytesConsumed, "")
		subpoints := ls.Subpoints()
		for j := 0; j < len(expectedSubpoints[i]); j++ {
			ep := expectedSubpoints[i][j]
			sp := subpoints[j]
			msg := fmt.Sprintf("In linestring test case %v/point %v, expected %v, got %v", i, j, ep, sp)
			assert.InDelta(t, ep.X(), sp.X(), 0.0000001, msg)
			assert.InDelta(t, ep.Y(), sp.Y(), 0.0000001, msg)
		}
		assert.Equal(t, expectedTypes[i], ls.Type(), "")
	}
}


type G interface {
	Area(x int, y int) int
}

type Gs struct {
	name string
}

func (gs *Gs) Init(name string) {
	gs.name = name
}

func (gs Gs) Area(x int, y int) int {
	return x*y
}
