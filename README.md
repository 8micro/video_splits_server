video path rules:

prefix/userid/GUID/ -->filename
                   +-->dirname(<==>filename)/xxx.m3u8
                                         +-->xx1.ts xx2.ts .... xxn.ts


for example:


                                                              +userid+yyyymmdd+timestamp+origname+.mp4
															  
data/video/upload/138483/9a3c83eb-2560-42df-b9d7-901b54b5161f/138483-20180101-1514798616-testvideo.mp4

data/video/upload/138483/9a3c83eb-2560-42df-b9d7-901b54b5161f/138483-20180101-1514798616-testvideo.mp4/138483-20180101-1514798616-testvideo.m3u8




data/video/upload/138483/9a3c83eb-2560-42df-b9d7-901b54b5161f/138483-20180101-1514798616-testvideo.mp4/138483-20180101-1514798616-testvideo-01.ts

data/video/upload/138483/9a3c83eb-2560-42df-b9d7-901b54b5161f/138483-20180101-1514798616-testvideo.mp4/138483-20180101-1514798616-testvideo-02.ts

data/video/upload/138483/9a3c83eb-2560-42df-b9d7-901b54b5161f/138483-20180101-1514798616-testvideo.mp4/138483-20180101-1514798616-testvideo-03.ts

data/video/upload/138483/9a3c83eb-2560-42df-b9d7-901b54b5161f/138483-20180101-1514798616-testvideo.mp4/138483-20180101-1514798616-testvideo-04.ts

data/video/upload/138483/9a3c83eb-2560-42df-b9d7-901b54b5161f/138483-20180101-1514798616-testvideo.mp4/138483-20180101-1514798616-testvideo-05.ts