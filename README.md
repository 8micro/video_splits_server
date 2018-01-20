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


format:
//get info from these lines
	/*
		streams.stream.0.index=0
		streams.stream.0.codec_name="h264"
		streams.stream.0.codec_long_name="H.264 / AVC / MPEG-4 AVC / MPEG-4 part 10"
		streams.stream.0.profile="Main"
		streams.stream.0.codec_type="video"
		streams.stream.0.codec_time_base="1/30"
		streams.stream.0.codec_tag_string="avc1"
		streams.stream.0.codec_tag="0x31637661"
		streams.stream.0.width=800
		streams.stream.0.height=600
		streams.stream.0.coded_width=800
		streams.stream.0.coded_height=600
		streams.stream.0.has_b_frames=2
		streams.stream.0.sample_aspect_ratio="0:1"
		streams.stream.0.display_aspect_ratio="0:1"
		streams.stream.0.pix_fmt="yuv420p"
		streams.stream.0.level=31
		streams.stream.0.color_range="unknown"
		streams.stream.0.color_space="unknown"
		streams.stream.0.color_transfer="unknown"
		streams.stream.0.color_primaries="unknown"
		streams.stream.0.chroma_location="left"
		streams.stream.0.field_order="unknown"
		streams.stream.0.timecode="N/A"
		streams.stream.0.refs=1
		streams.stream.0.is_avc="true"
		streams.stream.0.nal_length_size="4"
		streams.stream.0.id="N/A"
		streams.stream.0.r_frame_rate="15/1"
		streams.stream.0.avg_frame_rate="15/1"
		streams.stream.0.time_base="1/15"
		streams.stream.0.start_pts=0
		streams.stream.0.start_time="0.000000"
		streams.stream.0.duration_ts=8002
		streams.stream.0.duration="533.466667"
		streams.stream.0.bit_rate="499122"
		streams.stream.0.max_bit_rate="N/A"
		streams.stream.0.bits_per_raw_sample="8"
		streams.stream.0.nb_frames="8002"
		streams.stream.0.nb_read_frames="N/A"
		streams.stream.0.nb_read_packets="N/A"
		streams.stream.0.disposition.default=1
		streams.stream.0.disposition.dub=0
		streams.stream.0.disposition.original=0
		streams.stream.0.disposition.comment=0
		streams.stream.0.disposition.lyrics=0
		streams.stream.0.disposition.karaoke=0
		streams.stream.0.disposition.forced=0
		streams.stream.0.disposition.hearing_impaired=0
		streams.stream.0.disposition.visual_impaired=0
		streams.stream.0.disposition.clean_effects=0
		streams.stream.0.disposition.attached_pic=0
		streams.stream.0.disposition.timed_thumbnails=0
		streams.stream.0.tags.creation_time="1970-01-01T00:00:00.000000Z"
		streams.stream.0.tags.language="und"
		streams.stream.0.tags.handler_name="VideoHandler"
		streams.stream.1.index=1
		streams.stream.1.codec_name="aac"
		streams.stream.1.codec_long_name="AAC (Advanced Audio Coding)"
		streams.stream.1.profile="LC"
		streams.stream.1.codec_type="audio"
		streams.stream.1.codec_time_base="1/44100"
		streams.stream.1.codec_tag_string="mp4a"
		streams.stream.1.codec_tag="0x6134706d"
		streams.stream.1.sample_fmt="fltp"
		streams.stream.1.sample_rate="44100"
		streams.stream.1.channels=2
		streams.stream.1.channel_layout="stereo"
		streams.stream.1.bits_per_sample=0
		streams.stream.1.id="N/A"
		streams.stream.1.r_frame_rate="0/0"
		streams.stream.1.avg_frame_rate="0/0"
		streams.stream.1.time_base="1/44100"
		streams.stream.1.start_pts=0
		streams.stream.1.start_time="0.000000"
		streams.stream.1.duration_ts=23535641
		streams.stream.1.duration="533.688005"
		streams.stream.1.bit_rate="96000"
		streams.stream.1.max_bit_rate="96000"
		streams.stream.1.bits_per_raw_sample="N/A"
		streams.stream.1.nb_frames="22984"
		streams.stream.1.nb_read_frames="N/A"
		streams.stream.1.nb_read_packets="N/A"
		streams.stream.1.disposition.default=1
		streams.stream.1.disposition.dub=0
		streams.stream.1.disposition.original=0
		streams.stream.1.disposition.comment=0
		streams.stream.1.disposition.lyrics=0
		streams.stream.1.disposition.karaoke=0
		streams.stream.1.disposition.forced=0
		streams.stream.1.disposition.hearing_impaired=0
		streams.stream.1.disposition.visual_impaired=0
		streams.stream.1.disposition.clean_effects=0
		streams.stream.1.disposition.attached_pic=0
		streams.stream.1.disposition.timed_thumbnails=0
		streams.stream.1.tags.creation_time="1970-01-01T00:00:00.000000Z"
		streams.stream.1.tags.language="und"
		streams.stream.1.tags.handler_name="SoundHandler"
		format.filename="test.mp4"
		format.nb_streams=2
		format.nb_programs=0
		format.format_name="mov,mp4,m4a,3gp,3g2,mj2"
		format.format_long_name="QuickTime / MOV"
		format.start_time="0.000000"
		format.duration="533.688000"
		format.size="39930706"
		format.bit_rate="598562"
		format.probe_score=100
		format.tags.major_brand="isom"
		format.tags.minor_version="512"
		format.tags.compatible_brands="isomiso2avc1mp41"
		format.tags.creation_time="1970-01-01T00:00:00.000000Z"
		format.tags.encoder="Lavf53.24.2
	*/