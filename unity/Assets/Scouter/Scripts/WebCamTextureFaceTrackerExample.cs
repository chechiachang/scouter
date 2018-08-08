using UnityEngine;
using System;
using System.Collections;
using UnityEngine.Networking;
using System.IO;

using System.Collections.Generic;
using UnityEngine.UI;

#if UNITY_5_3 || UNITY_5_3_OR_NEWER
using UnityEngine.SceneManagement;
#endif
using OpenCVForUnity;
using OpenCVFaceTracker;

namespace FaceTrackerExample
{
    [Serializable]
    public class PostRequestBody
    {
      public string data;
    }
    
    [Serializable]
    public class User
    {
        public string login { get; set; }
        public int id { get; set; }
        public string avatarurl { get; set; }
        public string htmlurl { get; set; }
        public string gravatarid { get; set; }
        public string name { get; set; }
        public string company { get; set; }
        public string blog { get; set; }
        public string location { get; set; }
        public object email { get; set; }
        public object hireable { get; set; }
        public object bio { get; set; }
        public int publicrepos { get; set; }
        public int publicgists { get; set; }
        public int followers { get; set; }
        public int following { get; set; }
        public object createdat { get; set; }
        public object updatedat { get; set; }
        public object suspendedat { get; set; }
        public string type { get; set; }
        public bool siteadmin { get; set; }
        public object totalprivaterepos { get; set; }
        public object ownedprivaterepos { get; set; }
        public object privategists { get; set; }
        public object diskusage { get; set; }
        public object collaborators { get; set; }
        public object plan { get; set; }
        public string url     { get; set; }
        public string eventsurl { get; set; }
        public string followingurl { get; set; }
        public string followersurl { get; set; }
        public string gistsurl { get; set; }
        public string organizationsurl { get; set; }
        public string receivedeventsurl { get; set; }
        public string reposurl { get; set; }
        public string starredurl { get; set; }
        public string subscriptionsurl { get; set; }
        public List<object> textmatches { get; set; }
        public object permissions { get; set; }
    }
    
    [Serializable]
    public class ResponseObject
    {
        public int id;
        public int contribution;
        public int followers;
        public int publicgists;
        public int publicrepos;
    }

    /// <summary>
    /// WebCamTexture face tracker example.
    /// </summary>
    [RequireComponent(typeof(WebCamTextureToMatHelper))]
    public class WebCamTextureFaceTrackerExample : MonoBehaviour
    {

        /// <summary>
        /// The auto reset mode. if ture, Only if face is detected in each frame, face is tracked.
        /// </summary>
        public bool isAutoResetMode = true;

        public string apiserverip;

        /// <summary>
        /// The auto reset mode toggle.
        /// </summary>
        public Toggle isAutoResetModeToggle;
        
        /// <summary>
        /// The gray mat.
        /// </summary>
        Mat grayMat;

        public Texture2D m2Texture;
        public int id = 0;
        public int contribution = 0;
        public int publicrepos = 0;
        public int publicgists = 0;
        public int followers = 0;

        bool isRequestCompleted = true;

        public UnityEngine.UI.Text contributionText;
        public UnityEngine.UI.Text followersText;
        public UnityEngine.UI.Text reposText;
        public UnityEngine.UI.Text gistsText;

        
        /// <summary>
        /// The texture.
        /// </summary>
        Texture2D texture;
        
        /// <summary>
        /// The cascade.
        /// </summary>
        CascadeClassifier cascade;
        
        /// <summary>
        /// The face tracker.
        /// </summary>
        FaceTracker faceTracker;
        
        /// <summary>
        /// The face tracker parameters.
        /// </summary>
        FaceTrackerParams faceTrackerParams;

        /// <summary>
        /// The web cam texture to mat helper.
        /// </summary>
        WebCamTextureToMatHelper webCamTextureToMatHelper;

        /// <summary>
        /// The tracker_model_json_filepath.
        /// </summary>
        private string tracker_model_json_filepath;
        
        /// <summary>
        /// The haarcascade_frontalface_alt_xml_filepath.
        /// </summary>
        private string haarcascade_frontalface_alt_xml_filepath;


        // Use this for initialization
        void Start()
        {
            webCamTextureToMatHelper = gameObject.GetComponent<WebCamTextureToMatHelper>();

            isAutoResetModeToggle.isOn = isAutoResetMode;

            #if UNITY_WEBGL && !UNITY_EDITOR
            StartCoroutine(getFilePathCoroutine());
            #else
            tracker_model_json_filepath = Utils.getFilePath("tracker_model.json");
            haarcascade_frontalface_alt_xml_filepath = Utils.getFilePath("haarcascade_frontalface_alt.xml");
            Run();
            #endif
            
        }

        #if UNITY_WEBGL && !UNITY_EDITOR
        private IEnumerator getFilePathCoroutine()
        {
            var getFilePathAsync_0_Coroutine = StartCoroutine(Utils.getFilePathAsync("tracker_model.json", (result) => {
                tracker_model_json_filepath = result;
            }));
            var getFilePathAsync_1_Coroutine = StartCoroutine(Utils.getFilePathAsync("haarcascade_frontalface_alt.xml", (result) => {
                haarcascade_frontalface_alt_xml_filepath = result;
            }));
            
            
            yield return getFilePathAsync_0_Coroutine;
            yield return getFilePathAsync_1_Coroutine;
            
            Run();
        }
        #endif

        private void Run()
        {
            //initialize FaceTracker
            faceTracker = new FaceTracker(tracker_model_json_filepath);
            //initialize FaceTrackerParams
            faceTrackerParams = new FaceTrackerParams();

            cascade = new CascadeClassifier();
            cascade.load(haarcascade_frontalface_alt_xml_filepath);
//            if (cascade.empty())
//            {
//                Debug.LogError("cascade file is not loaded.Please copy from “FaceTrackerExample/StreamingAssets/” to “Assets/StreamingAssets/” folder. ");
//            }


            webCamTextureToMatHelper.Initialize();

        }

        /// <summary>
        /// Raises the webcam texture to mat helper initialized event.
        /// </summary>
        public void OnWebCamTextureToMatHelperInitialized()
        {
            Debug.Log("OnWebCamTextureToMatHelperInitialized");
            
            Mat webCamTextureMat = webCamTextureToMatHelper.GetMat();
            
            texture = new Texture2D(webCamTextureMat.cols(), webCamTextureMat.rows(), TextureFormat.RGBA32, false);


            gameObject.transform.localScale = new Vector3(webCamTextureMat.cols(), webCamTextureMat.rows(), 1);
            Debug.Log("Screen.width " + Screen.width + " Screen.height " + Screen.height + " Screen.orientation " + Screen.orientation);
            
            float width = 0;
            float height = 0;
            
            width = gameObject.transform.localScale.x;
            height = gameObject.transform.localScale.y;
            
            float widthScale = (float)Screen.width / width;
            float heightScale = (float)Screen.height / height;
            if (widthScale < heightScale)
            {
                Camera.main.orthographicSize = (width * (float)Screen.height / (float)Screen.width) / 2;
            } else
            {
                Camera.main.orthographicSize = height / 2;
            }
            
            gameObject.GetComponent<Renderer>().material.mainTexture = texture;



            grayMat = new Mat(webCamTextureMat.rows(), webCamTextureMat.cols(), CvType.CV_8UC1);

            
        }

        /// <summary>
        /// Raises the webcam texture to mat helper disposed event.
        /// </summary>
        public void OnWebCamTextureToMatHelperDisposed()
        {
            Debug.Log("OnWebCamTextureToMatHelperDisposed");

            faceTracker.reset();
            grayMat.Dispose();
        }

        /// <summary>
        /// Raises the webcam texture to mat helper error occurred event.
        /// </summary>
        /// <param name="errorCode">Error code.</param>
        public void OnWebCamTextureToMatHelperErrorOccurred(WebCamTextureToMatHelper.ErrorCode errorCode)
        {
            Debug.Log("OnWebCamTextureToMatHelperErrorOccurred " + errorCode);
        }
            
        // Update is called once per frame
        void Update()
        {

            if (webCamTextureToMatHelper.IsPlaying() && webCamTextureToMatHelper.DidUpdateThisFrame())
            {
                
                Mat rgbaMat = webCamTextureToMatHelper.GetMat();

                //convert image to greyscale
                Imgproc.cvtColor(rgbaMat, grayMat, Imgproc.COLOR_RGBA2GRAY);
                                        
                                            
                if (isAutoResetMode || faceTracker.getPoints().Count <= 0)
                {
//                                      Debug.Log ("detectFace");

                    //convert image to greyscale
                    using (Mat equalizeHistMat = new Mat())
                    using (MatOfRect faces = new MatOfRect())
                    {
                                                
                        Imgproc.equalizeHist(grayMat, equalizeHistMat);
                                                
                        cascade.detectMultiScale(equalizeHistMat, faces, 1.1f, 2, 0
                        //                                                                                 | Objdetect.CASCADE_FIND_BIGGEST_OBJECT
                        | Objdetect.CASCADE_SCALE_IMAGE, new OpenCVForUnity.Size(equalizeHistMat.cols() * 0.15, equalizeHistMat.cols() * 0.15), new Size());
                                                
                        if (faces.rows() > 0)
                        {
//                                              Debug.Log ("faces " + faces.dump ());

                            List<OpenCVForUnity.Rect> rectsList = faces.toList();
                            List<Point[]> pointsList = faceTracker.getPoints();

                            if (isAutoResetMode)
                            {
                                //add initial face points from MatOfRect
                                if (pointsList.Count <= 0)
                                {
                                    faceTracker.addPoints(faces);
//                                                                      Debug.Log ("reset faces ");
                                } else
                                {
                                                        
                                    // Disable nose rect display
                                    //for (int i = 0; i < rectsList.Count; i++)
                                    //{
                                    //    OpenCVForUnity.Rect trackRect = new OpenCVForUnity.Rect(rectsList [i].x + rectsList [i].width / 3, rectsList [i].y + rectsList [i].height / 2, rectsList [i].width / 3, rectsList [i].height / 3);
                                    //    //It determines whether nose point has been included in trackRect.                                      
                                    //    if (i < pointsList.Count && !trackRect.contains(pointsList [i] [67]))
                                    //    {
                                    //        rectsList.RemoveAt(i);
                                    //        pointsList.RemoveAt(i);
//                                  //                                                    Debug.Log ("remove " + i);
                                    //    }
                                    //    Imgproc.rectangle(rgbaMat, new Point(trackRect.x, trackRect.y), new Point(trackRect.x + trackRect.width, trackRect.y + trackRect.height), new Scalar(0, 0, 255, 255), 2);
                                    //}
                                }
                            } else
                            {
                                faceTracker.addPoints(faces);
                            }
                            //draw face rect with largest rect
                            int l = 0;
                            for (int i = 0; i < rectsList.Count; i++)
                            {
                                if (rectsList [i].area() > rectsList [l].area())
                                {
                                  l = i;
                                }

                                #if OPENCV_2
                                Core.rectangle (rgbaMat, new Point (rectsList [l].x, rectsList [l].y), new Point (rectsList [l].x + rectsList [l].width, rectsList [l].y + rectsList [l].height), new Scalar (255, 0, 0, 255), 2);
                                #else
                                Imgproc.rectangle(rgbaMat, new Point(rectsList [l].x, rectsList [l].y), new Point(rectsList [l].x + rectsList [l].width, rectsList [l].y + rectsList [l].height), new Scalar(255, 0, 0, 255), 2);
                                #endif
                                
                            }

                            //grayscale or rgba
                            //Mat croppedImage = new Mat(grayMat, rectsList[l]);

                            int x = rectsList[l].x;
                            int y = rectsList[l].y;
                            int w = rectsList[l].width;
                            int h = rectsList[l].height;

                            Color[] c = texture.GetPixels (x, y, w, h);
                            m2Texture = new Texture2D (w, h);
                            m2Texture.SetPixels (c);
                            m2Texture.Apply ();

                            byte[] imageBytes = m2Texture.EncodeToJPG();
                            Destroy(m2Texture);

                            //Debug
                            //File.WriteAllBytes(Application.dataPath + "/image.jpg", imageBytes);

                            //StartCoroutine(PostRequest("http://localhost:5000/face_detection", imageBytes));
                            if (isRequestCompleted){
                                StartCoroutine(PostRequest("http://"+apiserverip+":5000/face_detection", imageBytes));
                            }

                            // Display rect on texture
                                               
                        } else
                        {
                            if (isAutoResetMode)
                            {
                                faceTracker.reset();
                            }
                        }
                    }
                                            
                }

                //track face points.if face points <= 0, always return false.
                //if (faceTracker.track(grayMat, faceTrackerParams))
                //    faceTracker.draw(rgbaMat, new Scalar(255, 0, 0, 255), new Scalar(0, 255, 0, 255));
                                        
                //#if OPENCV_2
                //Core.putText (rgbaMat, "Contribution " + contribution + " repos " + publicRepos + " followers " + followers + " gists " + publicGists, new Point (5, rgbaMat.rows () - 5), Core.FONT_HERSHEY_SIMPLEX, 0.8, new Scalar (255, 255, 255, 255), 2, Core.LINE_AA, false);
                //#else
                //Imgproc.putText(rgbaMat, "Contribution " + contribution + " repos " + publicRepos + " followers " + followers + " gists " + publicGists, new Point(5, rgbaMat.rows() - 5), Core.FONT_HERSHEY_SIMPLEX, 0.8, new Scalar(255, 255, 255, 255), 2, Imgproc.LINE_AA, false);
                //#endif
                                        
                                        
                //Core.putText (rgbaMat, "W:" + rgbaMat.width () + " H:" + rgbaMat.height () + " SO:" + Screen.orientation, new Point (5, rgbaMat.rows () - 10), Core.FONT_HERSHEY_SIMPLEX, 1.0, new Scalar (255, 255, 255, 255), 2, Core.LINE_AA, false);

                Utils.matToTexture2D(rgbaMat, texture, webCamTextureToMatHelper.GetBufferColors());
                                        
            }
                                    
            if (Input.GetKeyUp(KeyCode.Space) || Input.touchCount > 0)
            {
                faceTracker.reset();
            }
                    
        }


        IEnumerator PostRequest(string url, byte[] bytes)
        {
            isRequestCompleted = false;

            PostRequestBody body = new PostRequestBody();
            body.data = Convert.ToBase64String(bytes);

            var uwr = new UnityWebRequest(url, "POST");
            var jsonBytes = System.Text.Encoding.UTF8.GetBytes(JsonUtility.ToJson(body));
            uwr.uploadHandler = new UploadHandlerRaw(jsonBytes);
            uwr.downloadHandler = new DownloadHandlerBuffer();
            uwr.SetRequestHeader("Content-Type", "application/json");
            //Send the request then wait here until it returns
            yield return uwr.SendWebRequest();
            if (uwr.isNetworkError)
            {
                Debug.Log("Error While Sending: " + uwr.error);
                isRequestCompleted = true;
            }
            else
            {
                //Debug.Log("Received: " + uwr.downloadHandler.text);
                
                ResponseObject obj = JsonUtility.FromJson<ResponseObject>(uwr.downloadHandler.text);
                //id = 0;
                //contribution = 0;
                //publicRepos = 0;
                //publicGist = 0;
                //followers = 0;
                //Debug.Log("Data: " + obj?.id + " "+ obj?.contribution + " " + obj?.publicrepos + " " + obj?.followers);
                if (obj?.contribution is int)
                {

                    contribution = obj.contribution;
                    id = obj.id;
                    publicrepos = obj.publicrepos;
                    publicgists = obj.publicgists;
                    followers = obj.followers;

                    // FIXME Cheating
                    if (contribution != 0)
                    {
                        contributionText.text = "Contribution: " + contribution;
                        followersText.text = "Followers: " + followers;
                        reposText.text = "Repos: " + publicrepos;
                        gistsText.text = "Gists: " + publicgists;
                    }
                }
            }

            isRequestCompleted = true;

            //UnityWebRequest uwr = new UnityWebRequest( url , UnityWebRequest.kHttpVerbPOST );
            //UploadHandlerRaw MyUploadHandler = new UploadHandlerRaw( bytes );
            //MyUploadHandler.contentType= "application/x-www-form-urlencoded"; // might work with 'multipart/form-data'
            //uwr.uploadHandler= MyUploadHandler;

            //Send the request then wait here until it returns
            //yield return uwr.SendWebRequest();
            //if (uwr.isNetworkError)
            //{
            //    Debug.Log("Error While Sending: " + uwr.error);
            //}
            //else
            //{
            //    Debug.Log("Received: " + uwr.downloadHandler.text);
            //}

            //WWWForm form = new WWWForm();
            //string bytestring = System.Text.Encoding.UTF8.GetString(bytes);
            //form.AddField("file", bytestring);

            //UnityWebRequest www = UnityWebRequest.Post(url, form);
            //www.SetRequestHeader("Content-Type", "multipart/form-data");
            //yield return www.SendWebRequest();

            //if (www.isNetworkError || www.isHttpError)
            //{
            //    Debug.Log(www.error);
            //}
            //else
            //{
            //    Debug.Log("Form upload complete!");
            //}
        }

        /// <summary>
        /// Raises the disable event.
        /// </summary>
        void OnDisable()
        {
            webCamTextureToMatHelper.Dispose();

            if (cascade != null)
                cascade.Dispose();
        }

        /// <summary>
        /// Raises the back button event.
        /// </summary>
        public void OnBackButton()
        {
            #if UNITY_5_3 || UNITY_5_3_OR_NEWER
            SceneManager.LoadScene("FaceTrackerExample");
            #else
            Application.LoadLevel("FaceTrackerExample");
            #endif
        }

        public void OnApiserverIpchange(string ip)
        {
            apiserverip = ip;
            Debug.Log("apiserverip set to " + apiserverip);
        }

        /// <summary>
        /// Raises the play button event.
        /// </summary>
        public void OnPlayButton()
        {
            webCamTextureToMatHelper.Play();
        }

        /// <summary>
        /// Raises the pause button event.
        /// </summary>
        public void OnPauseButton()
        {
            webCamTextureToMatHelper.Pause();
        }

        /// <summary>
        /// Raises the stop button event.
        /// </summary>
        public void OnStopButton()
        {
            webCamTextureToMatHelper.Stop();
        }

        /// <summary>
        /// Raises the change camera button event.
        /// </summary>
        public void OnChangeCameraButton()
        {
            webCamTextureToMatHelper.Initialize(null, webCamTextureToMatHelper.requestedWidth, webCamTextureToMatHelper.requestedHeight, !webCamTextureToMatHelper.requestedIsFrontFacing);
        }

        /// <summary>
        /// Raises the change auto reset mode toggle event.
        /// </summary>
        public void OnIsAutoResetModeToggle()
        {
            if (isAutoResetModeToggle.isOn)
            {
                isAutoResetMode = true;
            } else
            {
                isAutoResetMode = false;
            }
        }
                
    }
}
