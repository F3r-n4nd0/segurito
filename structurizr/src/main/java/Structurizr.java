import com.structurizr.Workspace;
import com.structurizr.api.StructurizrClient;
import com.structurizr.api.StructurizrClientException;
import com.structurizr.documentation.StructurizrDocumentationTemplate;
import com.structurizr.model.*;
import com.structurizr.view.*;

import java.io.File;

public class Structurizr {

    private static final long WORKSPACE_ID = Long.parseLong(System.getenv("WORKSPACE_ID"));
    private static final String API_KEY = System.getenv("API_KEY");
    private static final String API_SECRET = System.getenv("API_SECRET");

    private static final String EXISTING_SYSTEM_TAG = "Existing System";
    private static final String WEB_BROWSER_TAG = "Web Browser";
    private static final String DATABASE_TAG = "Database";
    private static final String PIPE_TAG = "Pipe";
    private static final String MICROSERVICE_TAG = "Microservice";


    public static void main(String[] args) throws Exception {

        Workspace workspace = new Workspace("SEGURITO", "Software de control de asistencia.");
        Model model = workspace.getModel();
        ViewSet views = workspace.getViews();

        //CONTEXT

        Person staff = model.addPerson("Personal", "Una persona empleada por la organización..");
        Person restfulClient = model.addPerson("Cliente Restful", "Cliente lector de eventos.");

        SoftwareSystem seguritoAttendanceMonitoringSystem = model.addSoftwareSystem("SEGURITO", "Software de control de asistencia");
        staff.uses(seguritoAttendanceMonitoringSystem, "Registra su entrada/salida","HTTPS");
        restfulClient.uses(seguritoAttendanceMonitoringSystem, "Consulta eventos por usuario","RESTFUL/JSON");

        SoftwareSystem mesaLegacyERPSystem = model.addSoftwareSystem("MESA", "Sistema ERP legacy");
        mesaLegacyERPSystem.addTags(EXISTING_SYSTEM_TAG);
        seguritoAttendanceMonitoringSystem.uses(mesaLegacyERPSystem, "Autentificacion","STDIN/STDOUT");

        SoftwareSystem pumariLogSystem = model.addSoftwareSystem("PUMARI", "Sistema de registro de eventos");
        pumariLogSystem.addTags(EXISTING_SYSTEM_TAG);
        seguritoAttendanceMonitoringSystem.uses(pumariLogSystem, "Registra entrada/salida", "STDIN/STDOUT");


        SystemContextView contextView = views.createSystemContextView(seguritoAttendanceMonitoringSystem, "Contexto", "Diagrama de contexto del sistema para la arquitectura del sistema de monitoreo de asistencia.");
        contextView.addAllElements();


        //CONTAINER

        Container evoAccessSystem = seguritoAttendanceMonitoringSystem.addContainer("EVO", "Modulo de acceso web.", "Java Script Angular");
        evoAccessSystem.addTags(WEB_BROWSER_TAG);
        staff.uses(evoAccessSystem, "Registra entrada/salida");

        Container apiApplication = seguritoAttendanceMonitoringSystem.addContainer("Control de asistencia API", "Provee la funcionalidad para el control de asistencia.","Go");
        evoAccessSystem.uses(apiApplication, "Hace llamadas a la API", "jSON/HTTPS");
        apiApplication.uses(mesaLegacyERPSystem, "Autentifica","STDIN/STDOUT");


        Container camachoEventAPI = seguritoAttendanceMonitoringSystem.addContainer("CAMACHO API", "Modulo de consulta de eventos.","Go");
        restfulClient.uses(camachoEventAPI,"Consulta eventos por usuario", "jSON/HTTPS");

        Container messageBus = seguritoAttendanceMonitoringSystem.addContainer("Bus de mensajes", "Mensajes de registro de eventos.", "rabbitmq");
        messageBus.addTags(PIPE_TAG);
        apiApplication.uses(messageBus, "Registra entrada/salida");

        Container eventService = seguritoAttendanceMonitoringSystem.addContainer("Servicio registro eventos", "Provee la funcionalidad para el registor de eventos.","Go");
        eventService.addTags(MICROSERVICE_TAG);
        eventService.uses(pumariLogSystem, "Registra entrada/salida","STDIN/STDOUT");
        eventService.uses(messageBus,"Lee los eventos");

        Container eventsDataBase = seguritoAttendanceMonitoringSystem.addContainer("Base de datos", "Base de datos que registra los eventos de entradas/salidas","");
        eventsDataBase.addTags(DATABASE_TAG);
        camachoEventAPI.uses(eventsDataBase, "Lee eventos por usuario");
        apiApplication.uses(eventsDataBase, "Almacena entrada/salida");

        ContainerView containerView = views.createContainerView(seguritoAttendanceMonitoringSystem, "Contenedores", "El diagrama de contenedores para el sistema de control de asistencia.");
        containerView.addAllContainersAndInfluencers();
        containerView.setPaperSize(PaperSize.A4_Landscape);


        //COMPONENTES

            //API

        Component registerEvents = apiApplication.addComponent("Registrar eventos", "Registra los venetos de entrada/salida a un sistema externo.", "Go");
        registerEvents.uses(messageBus, "Usa");

        Component storeEvents = apiApplication.addComponent("Almacena eventos", "Almacena los venetos de entrada/salida", "Go");
        storeEvents.uses(eventsDataBase, "Usa");

        Component authentication = apiApplication.addComponent("Autentificacion", "Autentifica un usaurio a un sistema externo.", "Go");
        authentication.uses(mesaLegacyERPSystem, "Usa");

        Component assistanceController = apiApplication.addComponent("Control de asistencia", "Permite hacer el control de asistencia del personal.", "Go");
        assistanceController.uses(storeEvents, "Usa");
        assistanceController.uses(registerEvents, "Usa");
        assistanceController.uses(authentication, "Usa");
        evoAccessSystem.uses(assistanceController, "Usa");

        ComponentView componentView = views.createComponentView(apiApplication, "Componentes - API", "El diagrama de componentes.");
        componentView.addAllComponents();
        componentView.addExternalDependencies();

            //CAMACHO API

        Component eventReader = camachoEventAPI.addComponent("Lector de eventos", "Lector los venetos de entrada/salida.", "Go");
        eventReader.uses(eventsDataBase, "Usa");

        Component eventQuery = camachoEventAPI.addComponent("Consulta eventos", "Consulta los venetos de entrada/salida por usuario.", "Go");
        eventQuery.uses(eventReader, "Usa");
        restfulClient.uses(eventQuery, "Usa");

        ComponentView componentViewCamacho = views.createComponentView(camachoEventAPI, "Componentes - CAMACHO API", "El diagrama de componentes.");
        componentViewCamacho.addAllComponents();
        componentViewCamacho.addExternalDependencies();

        //Style

        setStyle(views);

        StructurizrDocumentationTemplate template = new StructurizrDocumentationTemplate(workspace);
        File documentationRoot = new File("./documentation");
        template.addContextSection(seguritoAttendanceMonitoringSystem, new File(documentationRoot, "context.md"));
//        template.addFunctionalOverviewSection(mesaLegacyERPSystem, new File(documentationRoot, "functional-overview.md"));
        template.addImages(documentationRoot);

        setWorkspaceToStructurizr(workspace);


    }

    private static void setStyle(ViewSet views) {
        Styles styles = views.getConfiguration().getStyles();
        styles.addElementStyle(Tags.ELEMENT).color("#ffffff");
        styles.addElementStyle(Tags.SOFTWARE_SYSTEM).background("#1168bd");
        styles.addElementStyle(Tags.CONTAINER).background("#438dd5");
        styles.addElementStyle(Tags.COMPONENT).background("#85bbf0").color("#000000");
        styles.addElementStyle(Tags.PERSON).background("#08427b").shape(Shape.Person).fontSize(22);
        styles.addElementStyle(EXISTING_SYSTEM_TAG).background("#999999");
        styles.addElementStyle(WEB_BROWSER_TAG).shape(Shape.WebBrowser);
        styles.addElementStyle(DATABASE_TAG).shape(Shape.Cylinder);
        styles.addElementStyle(PIPE_TAG).shape(Shape.Pipe);
        styles.addElementStyle(MICROSERVICE_TAG).shape(Shape.Hexagon);

    }

    private static void setWorkspaceToStructurizr(Workspace workspace) throws StructurizrClientException {
        StructurizrClient structurizrClient = new StructurizrClient(API_KEY, API_SECRET);
        structurizrClient.putWorkspace(WORKSPACE_ID, workspace);
    }

}